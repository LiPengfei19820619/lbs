package log

import (
	"log"
	"os"
)

func newFileLogger(fileName string) *log.Logger {
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil
	}

	fileLogger := log.New(logFile, "lbs:", log.Ldate|log.Lmicroseconds|log.Lshortfile)

	return fileLogger
}
