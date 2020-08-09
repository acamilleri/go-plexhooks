package plexhooks

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"github.com/acamilleri/go-plexhooks/plex"
)

// Definition app configuration
type Definition struct {
	ListenAddr *net.TCPAddr
	Actions    *Actions
	Logger     LoggerDefinition
}

// App struct with listenAddr and actions
type App struct {
	listenAddr *net.TCPAddr
	actions    *Actions
	log        *logrus.Logger
}

// New create app from Definition values
func New(definition Definition) *App {
	log := setupLogger(definition.Logger)

	return &App{
		listenAddr: definition.ListenAddr,
		actions:    definition.Actions,
		log:        log,
	}
}

// Run running app with http server
func (a *App) Run() error {
	if a.log.GetLevel() == logrus.DebugLevel {
		for hook, actions := range *a.actions {
			for _, action := range actions {
				a.log.Debugf("action %s registered for hook %s", action.Name(), hook)
			}
		}
	}

	http.HandleFunc("/events", a.handler())
	http.Handle("/metrics", promhttp.Handler())

	a.log.Infof("running server on %s", a.listenAddr)
	return http.ListenAndServe(a.listenAddr.String(), nil)
}

func (a *App) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.WriteHeader(http.StatusOK)

		event, err := parseRequest(r)
		if err != nil {
			a.log.WithError(err).Error("failed to parse request")
		}

		a.log.Infof("%s event handled", event.Name)
		err = a.triggerActionsOnEvent(event)
		if err != nil {
			a.log.WithError(err).Errorf("%s event actions failed", event.Name)
		}
	}
}

func (a *App) triggerActionsOnEvent(event plex.Event) error {
	eventsReceivedTotal.With(prometheus.Labels{"event": event.Name.String()}).Inc()
	hookName := event.Name

	actions := a.actions.GetByHook(hookName)
	if len(actions) == 0 {
		return fmt.Errorf("no actions registered for %s hook", hookName)
	}

	for _, action := range actions {
		name := action.Name()
		actionDuration := newTrackActionDuration(event, action)

		a.log.Debugf("action %s triggered", name)
		err := action.Execute(event)
		if err != nil {
			a.log.WithError(err).Errorf("action %s failed", name)
			actionsErrorTotal.With(
				prometheus.Labels{"event": event.Name.String(), "action": action.Name()},
			).Inc()
			actionDuration.Finish()
			continue
		}

		a.log.Infof("action %s success", name)
		actionDuration.Finish()
		actionsSuccessTotal.With(
			prometheus.Labels{"event": event.Name.String(), "action": action.Name()},
		).Inc()
	}

	return nil
}

func parseRequest(r *http.Request) (plex.Event, error) {
	err := r.ParseMultipartForm(64)
	if err != nil {
		return plex.Event{}, fmt.Errorf("parse multipart request failed: %v", err)
	}

	payloadReq, ok := r.MultipartForm.Value["payload"]
	if !ok {
		return plex.Event{}, fmt.Errorf("read payload failed: %v", err)
	}

	var event plex.Event
	err = json.Unmarshal([]byte(payloadReq[0]), &event)
	if err != nil {
		return plex.Event{}, fmt.Errorf("decode payload failed: %v", err)
	}

	return event, nil
}
