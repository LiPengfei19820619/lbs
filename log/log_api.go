package log

import (
	"log"
)

const (
	runLogFileName = "./lbs_run_log.txt"
)

var (
	runLogger *log.Logger
)

func init() {
	runLogger = newFileLogger(runLogFileName)
}

// WriteRunLog 记录运行日志
func WriteRunLog(v ...interface{}) {
	runLogger.Println(v)
}
