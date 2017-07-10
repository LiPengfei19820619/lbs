package log

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const (
	runLogFileName = "./lbs_run_log.txt"

	// LOG_ERROR 日志级别
	LOG_ERROR = iota
	LOG_WARNING
	LOG_DEBUG
)

const (
	// HTTPRoleServer HTTP角色为服务器
	HTTPRoleServer = iota
	// HTTPRoleClient HTTP角色为客户端
	HTTPRoleClient
)

/****** Variables ******/
var (
	// LogBufferLength specifies how many log messages a particular log4go
	// logger can buffer at a time before writing them.
	LogBufferLength = 32
)

// HTTPLogRecord HTTP日志信息
type HTTPLogRecord struct {
	Role       int       `json:"Role"`
	ClientAddr string    `json:"ClientAddr"`
	ServerAddr string    `json:"ServerAddr"`
	ReqTime    time.Time `json:"ReqTime"`
	ResTime    time.Time `json:"ResTime"`
	Request    string    `json:"Request"`
	Response   string    `json:"Response"`
}

var (
	runLogger  *log.Logger
	httpLogger *log.Logger
)

func init() {
	runLogger = newFileLogger(runLogFileName)
	httpLogger = newSocketLogger("udp", "127.0.0.1:12124")
}

// WriteRunLog 记录运行日志
func WriteRunLog(v ...interface{}) {
	runLogger.Println(v)
}

// WriteHTTPLog 记录HTTP消息日志
func WriteHTTPLog(logRecord *HTTPLogRecord) {
	log, _ := json.Marshal(logRecord)
	fmt.Println(string(log))
	httpLogger.Println(string(log))
}
