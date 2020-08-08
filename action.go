package plexhooks

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/acamilleri/go-plexhooks/plex"
)

// Action generic interface to create action
type Action interface {
	Name() string
	Execute(event plex.Event) error
}

// Actions list of Action by Event name
type Actions map[plex.Name][]Action

// NewActions return an initialize map to
// define actions to run on each event
func NewActions() *Actions {
	actions := make(Actions)
	return &actions
}

// Add define one action or multiple actions to run on Event
func (a *Actions) Add(hook plex.Name, actions ...Action) {
	actionsList := *a
	actionsList[hook] = append(actionsList[hook], actions...)
	*a = actionsList
}

func (a Actions) triggerActionsOnEvent(event plex.Event) error {
	hookName := event.Name

	if actions, ok := a[hookName]; ok {
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

	return fmt.Errorf("no actions registered for %s hook", hookName)
}
