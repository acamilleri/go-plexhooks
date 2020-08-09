package plexhooks

import "github.com/sirupsen/logrus"

// LoggerDefinition logger configuration
type LoggerDefinition struct {
	Level     logrus.Level
	Formatter logrus.Formatter
}

func setupLogger(definition LoggerDefinition) *logrus.Logger {
	var level logrus.Level = logrus.InfoLevel
	var formatter logrus.Formatter = &logrus.TextFormatter{}

	if definition != (LoggerDefinition{}) {
		level = definition.Level
		formatter = definition.Formatter
	}

	log := logrus.New()
	log.SetLevel(level)
	log.SetFormatter(formatter)
	return log
}
