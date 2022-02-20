package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LoggingOutput() {
	logFile, _ := os.Create("logs/api_requests.log")

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(logFile)
}
