package log

import (
	"log"
	"os"
)

func newFileLogger(fileName string) *log.Logger {
	if !fileExist(fileName) {
		os.Create(fileName)
	}

	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil
	}

	fileLogger := log.New(logFile, "lbs", 0)

	return fileLogger
}

func fileExist(filePath string) bool {
	_, err := os.Stat(filePath)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
