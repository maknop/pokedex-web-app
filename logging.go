package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func loggingOutput() {
	logFile, _ := os.Create("logs/api_requests.log")

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(logFile)
}
