package plexhooks

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

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
}

// New create app from Definition values
func New(definition Definition) *App {
	setupLogger(definition.Logger)

	return &App{
		listenAddr: definition.ListenAddr,
		actions:    definition.Actions,
	}
}

// Run running app with http server
func (a *App) Run() error {
	if logrus.GetLevel() == logrus.DebugLevel {
		for hook, actions := range *a.actions {
			for _, action := range actions {
				logrus.Debugf("action %s registered for hook %s", action.Name(), hook)
			}
		}
	}

	http.HandleFunc("/events", a.handler())

	logrus.Infof("running server on %s", a.listenAddr)
	return http.ListenAndServe(a.listenAddr.String(), nil)
}

func (a *App) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.WriteHeader(http.StatusOK)

		event, err := parseRequest(r)
		if err != nil {
			logrus.WithError(err).Error("failed to parse request")
		}

		logrus.Infof("%s event handled", event.Name)
		err = a.triggerActionsOnEvent(event)
		if err != nil {
			logrus.WithError(err).Errorf("%s event actions failed", event.Name)
		}
	}
}

func (a *App) triggerActionsOnEvent(event plex.Event) error {
	hookName := event.Name

	actions := a.actions.GetByHook(hookName)
	if len(actions) == 0 {
		return fmt.Errorf("no actions registered for %s hook", hookName)
	}

	for _, action := range actions {
		name := action.Name()

		logrus.Debugf("action %s triggered", name)
		err := action.Execute(event)
		if err != nil {
			logrus.WithError(err).Errorf("action %s failed", name)
			continue
		}
		logrus.Infof("action %s success", name)
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
