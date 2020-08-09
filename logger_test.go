package plexhooks

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_setupLogger(t *testing.T) {
	type args struct {
		definition LoggerDefinition
	}
	tests := []struct {
		name              string
		args              args
		expectedLevel     logrus.Level
		expectedFormatter logrus.Formatter
	}{
		{
			name: "SetupLoggerDefaultValues",
			args: args{
				definition: LoggerDefinition{},
			},
			expectedFormatter: &logrus.TextFormatter{},
			expectedLevel:     logrus.InfoLevel,
		},
		{
			name: "SetupLoggerWithDebugLevel",
			args: args{
				definition: LoggerDefinition{
					Level:     logrus.DebugLevel,
					Formatter: &logrus.TextFormatter{},
				},
			},
			expectedFormatter: &logrus.TextFormatter{},
			expectedLevel:     logrus.DebugLevel,
		},
		{
			name: "SetupLoggerWithJSONFormatter",
			args: args{
				definition: LoggerDefinition{
					Level:     logrus.DebugLevel,
					Formatter: &logrus.JSONFormatter{},
				},
			},
			expectedFormatter: &logrus.JSONFormatter{},
			expectedLevel:     logrus.DebugLevel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := setupLogger(tt.args.definition)

			assert.Equal(t, tt.expectedLevel, log.Level)
			assert.Equal(t, tt.expectedFormatter, log.Formatter)
		})
	}
}
