package plexhooks

import (
	"net"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	listenAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	loggerDefinition := LoggerDefinition{
		Level:     logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{},
	}

	app := New(Definition{
		ListenAddr: listenAddr,
		Actions:    nil,
		Logger:     loggerDefinition,
	})

	expectedApp := &App{
		listenAddr: listenAddr,
		actions:    nil,
		log:        app.log,
	}

	assert.Equal(t, expectedApp, app)
}
