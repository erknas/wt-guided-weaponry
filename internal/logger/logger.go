package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetupLogger() *log.Logger {
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}

	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.InfoLevel)

	return logger
}
