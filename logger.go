package plexhooks

import "github.com/sirupsen/logrus"

// LoggerDefinition logger configuration
type LoggerDefinition struct {
	Level     logrus.Level
	Formatter logrus.Formatter
}

func setupLogger(definition LoggerDefinition) {
	var level logrus.Level = logrus.InfoLevel
	var formatter logrus.Formatter = &logrus.TextFormatter{}

	if definition != (LoggerDefinition{}) {
		level = definition.Level
		formatter = definition.Formatter
	}

	logrus.SetLevel(level)
	logrus.SetFormatter(formatter)
}
