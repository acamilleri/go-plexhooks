package plexhooks

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/acamilleri/go-plexhooks/plex"
)

var (
	eventsReceivedTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "plexhooks",
			Subsystem: "events",
			Name:      "received_total",
			Help:      "Total number of events received.",
		}, []string{"event"})

	actionsSuccessTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "plexhooks",
			Subsystem: "actions",
			Name:      "success_total",
			Help:      "Total number of actions by hook executed with success.",
		}, []string{"event", "action"})

	actionsDurationDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "plexhooks",
			Subsystem: "actions",
			Name:      "duration_seconds",
			Help:      "A histogram of time to running action by hook.",
			Buckets:   prometheus.DefBuckets,
		}, []string{"event", "action"})

	actionsErrorTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "plexhooks",
			Subsystem: "actions",
			Name:      "error_total",
			Help:      "Total number of actions by hook failed.",
		}, []string{"event", "action"})
)

type trackActionDuration struct {
	event  string
	action string
	start  time.Time
	stop   time.Time
}

func newTrackActionDuration(event plex.Event, action Action) *trackActionDuration {
	return &trackActionDuration{
		event:  event.Name.String(),
		action: action.Name(),
		start:  time.Now(),
		stop:   time.Time{},
	}
}

func (track *trackActionDuration) Finish() {
	track.stop = time.Now()

	duration := track.stop.Sub(track.start).Seconds()
	actionsDurationDuration.With(prometheus.Labels{
		"event":  track.event,
		"action": track.action,
	}).Observe(duration)
}
