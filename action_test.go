package plexhooks

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/acamilleri/go-plexhooks/plex"
)

type myTestActionOnMediaPlay struct{}

func (m myTestActionOnMediaPlay) Name() string {
	return "myTestActionOnMediaPlay"
}

func (m myTestActionOnMediaPlay) Execute(event plex.Event) error {
	return nil
}

type myTest2ActionOnMediaPlay struct{}

func (m myTest2ActionOnMediaPlay) Name() string {
	return "myTest2ActionOnMediaPlay"
}

func (m myTest2ActionOnMediaPlay) Execute(event plex.Event) error {
	return nil
}

type myTestActionOnMediaResume struct{}

func (m myTestActionOnMediaResume) Name() string {
	return "myTestActionOnMediaResume"
}

func (m myTestActionOnMediaResume) Execute(event plex.Event) error {
	return nil
}

func TestActions_Add(t *testing.T) {
	type args struct {
		hookActions map[plex.Name][]Action
	}
	tests := []struct {
		name            string
		args            args
		expectedActions *Actions
	}{
		{
			name: "AddActionOnMediaPlay",
			args: args{
				hookActions: map[plex.Name][]Action{
					plex.MediaPlay: []Action{
						&myTestActionOnMediaPlay{},
					},
				},
			},
			expectedActions: &Actions{
				plex.MediaPlay: []Action{
					&myTestActionOnMediaPlay{},
				},
			},
		},
		{
			name: "NoActionAdded",
			args: args{
				hookActions: map[plex.Name][]Action{
					plex.MediaPlay: nil,
				},
			},
			expectedActions: &Actions{
				plex.MediaPlay: nil,
			},
		},
		{
			name: "AddMultipleActionOnMediaPlay",
			args: args{
				hookActions: map[plex.Name][]Action{
					plex.MediaPlay: []Action{
						&myTestActionOnMediaPlay{},
						&myTest2ActionOnMediaPlay{},
					},
				},
			},
			expectedActions: &Actions{
				plex.MediaPlay: []Action{
					&myTestActionOnMediaPlay{},
					&myTest2ActionOnMediaPlay{},
				},
			},
		},
		{
			name: "AddOneActionOnMultipleHook",
			args: args{
				hookActions: map[plex.Name][]Action{
					plex.MediaPlay: []Action{
						&myTestActionOnMediaPlay{},
						&myTest2ActionOnMediaPlay{},
					},
					plex.MediaResume: []Action{
						&myTestActionOnMediaResume{},
					},
				},
			},
			expectedActions: &Actions{
				plex.MediaPlay: []Action{
					&myTestActionOnMediaPlay{},
					&myTest2ActionOnMediaPlay{},
				},
				plex.MediaResume: []Action{
					&myTestActionOnMediaResume{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actions := NewActions()
			for hook, as := range tt.args.hookActions {
				actions.Add(hook, as...)
			}

			assert.Equal(t, tt.expectedActions, actions)
		})
	}
}

func TestActions_GetByHook(t *testing.T) {
	type args struct {
		hook plex.Name
	}
	tests := []struct {
		name    string
		actions Actions
		args    args
		want    []Action
	}{
		{
			name: "GetMediaPlayAction",
			args: args{hook: plex.MediaPlay},
			actions: Actions{
				plex.MediaPlay: []Action{
					&myTestActionOnMediaPlay{},
				},
			},
			want: []Action{
				&myTestActionOnMediaPlay{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.actions.GetByHook(tt.args.hook); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewActions(t *testing.T) {
	expectedActions := make(Actions)
	actions := NewActions()
	assert.Equal(t, &expectedActions, actions)
}
