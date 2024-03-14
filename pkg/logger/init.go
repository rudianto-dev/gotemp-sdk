package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	Logrus *logrus.Logger
}

func NewLogger(logrus *logrus.Logger) *Logger {
	return &Logger{
		Logrus: logrus,
	}
}
