package log

import (
	"fmt"
	"log"
	"time"
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
	dbLogger *log.Logger
)

func init() {
	dbLogger = NewDBLogger("D:\\lbs_http_log.sl3")
}

// WriteDBLog 记录运行日志
func WriteDBLog(v ...interface{}) {
	fmt.Println(v...)
	dbLogger.Println(v...)
}
