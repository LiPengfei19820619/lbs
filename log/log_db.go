package log

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

// DBLogWriter 将日志发送到数据库
type DBLogWriter struct {
	logs chan []byte
}

func NewDBLogger(dbfile string) *log.Logger {
	dbWriter := NewDBLogWriter(dbfile)
	if dbWriter == nil {
		return nil
	}
	dbLogger := log.New(dbWriter, "", 0)

	return dbLogger
}

func (writer *DBLogWriter) Write(log []byte) (int, error) {
	writer.logs <- log

	return len(log), nil
}

func (writer *DBLogWriter) Close() {
	close(writer.logs)
}

func NewDBLogWriter(dbfile string) *SocketLogWriter {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewDBLogWriter(%q): %s\n", dbfile, err)
		return nil
	}

	w := SocketLogWriter{make(chan []byte, LogBufferLength)}

	go func() {
		defer func() {
			if db != nil {
				db.Close()
			}
		}()

		for rec := range w.logs {
			err = writeHTTPLogToDB(db, rec)
			if err != nil {
				fmt.Fprintln(os.Stderr, "writeHTTPLogToDB: ", err.Error())
				return
			}
		}
	}()

	return &w
}

func writeHTTPLogToDB(db *sql.DB, log []byte) error {
	var httpLogRecord HTTPLogRecord

	err := json.Unmarshal(log, &httpLogRecord)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unmarshal http log json error: ", err.Error())
		return err
	}

	sql := fmt.Sprint("INSERT INTO http_log VALUES(",
		httpLogRecord.Role, ",",
		"'", httpLogRecord.ClientAddr, "',",
		"'", httpLogRecord.ServerAddr, "',",
		"'", httpLogRecord.ReqTime, "',",
		"'", httpLogRecord.ResTime, "',",
		"'", httpLogRecord.Request, "',",
		"'", httpLogRecord.Response, "')")

	_, err = db.Exec(sql)
	if err != nil {
		fmt.Fprintln(os.Stderr, "exec sql error:", err.Error(), ",sql:", sql)
		return err
	}

	return nil
}
