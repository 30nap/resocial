package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
}

func Info(msg string, args ...interface{}) {
	Log.Infof(msg, args...)
}

func Error(msg string, args ...interface{}) {
	Log.Errorf(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	Log.Warnf(msg, args...)
}
