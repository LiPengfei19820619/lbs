package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"zte/ims/lbs/log"
)

var (
	port = flag.String("p", "12124", "Port number to listen on")
)

func e(err error) {
	if err != nil {
		fmt.Printf("Erroring out: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	// Bind to the port
	bind, err := net.ResolveUDPAddr("udp", "0.0.0.0:"+*port)
	e(err)

	// Create listener
	listener, err := net.ListenUDP("udp", bind)
	e(err)

	fmt.Printf("Listening to port %s...\n", *port)
	for {
		// read into a new buffer
		buffer := make([]byte, 1024)

		len, _, err := listener.ReadFrom(buffer)
		e(err)

		// log to standard output
		fmt.Println(string(buffer))
		js := string(buffer[:len])
		logger := log.NewDBLogger("D:\\lbs_http_log.sl3")

		fmt.Println(js)
		logger.Println(js)
	}
}
