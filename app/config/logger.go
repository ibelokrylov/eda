package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
	logger.SetOutput(os.Stdout)
	return logger
}

var Logger = SetupLogger()
