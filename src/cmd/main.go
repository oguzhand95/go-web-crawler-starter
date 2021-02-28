package main

import (
	"fmt"
	"github.com/google/logger"
	"os"
)

const (
	logFileName = "go-web-crawler-starter.log"
)

func main() {
	defer logger.Init(logFileName, true, true, getLogFile()).Close()
}

func getLogFile() *os.File {
	logFile, err := os.OpenFile(fmt.Sprintf("./%s", logFileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}

	return logFile
}