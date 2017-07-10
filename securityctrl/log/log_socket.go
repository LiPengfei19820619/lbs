package log

import (
	"fmt"
	"log"
	"net"
	"os"
)

// SocketLogWriter 将日志发送到Socket
type SocketLogWriter struct {
	logs chan []byte
}

func newSocketLogger(proto, hostport string) *log.Logger {
	socketWriter := NewSocketLogWriter(proto, hostport)
	if socketWriter == nil {
		return nil
	}
	socketLogger := log.New(socketWriter, "", 0)

	return socketLogger
}

func (writer *SocketLogWriter) Write(log []byte) (int, error) {
	writer.logs <- log

	return len(log), nil
}

func (writer *SocketLogWriter) Close() {
	close(writer.logs)
}

func NewSocketLogWriter(proto, hostport string) *SocketLogWriter {
	sock, err := net.Dial(proto, hostport)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewSocketLogWriter(%q): %s\n", hostport, err)
		return nil
	}

	w := SocketLogWriter{make(chan []byte, LogBufferLength)}

	go func() {
		defer func() {
			if sock != nil && proto == "tcp" {
				sock.Close()
			}
		}()

		for rec := range w.logs {

			_, err = sock.Write(rec)
			if err != nil {
				fmt.Fprintf(os.Stderr, "SocketLogWriter(%q): %s", hostport, err)
				return
			}
		}
	}()

	return &w
}
