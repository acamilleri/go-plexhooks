package plexhooks

import (
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

// GetByHook get actions by hook name
func (a *Actions) GetByHook(hook plex.Name) []Action {
	actionsList := *a
	if actions, ok := actionsList[hook]; ok {
		return actions
	}
	return nil
}
