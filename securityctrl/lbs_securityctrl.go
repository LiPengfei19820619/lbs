package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/nakagami/firebirdsql"

	"time"
	"zte/ims/lbs/securityctrl/bll"
	"zte/ims/lbs/securityctrl/log"
)

func main() {
	fmt.Println("lbs start running ... ")
	log.WriteRunLog("aaa")

	httpLog := log.HTTPLogRecord{Role: 0}
	httpLog.ClientAddr = "127.0.0.1:1000"
	httpLog.ServerAddr = "127.0.0.1:3001"
	httpLog.ReqTime = time.Now()
	httpLog.ResTime = time.Now()
	httpLog.Request = "POST"
	httpLog.Response = "200"

	log.WriteHTTPLog(&httpLog)

	go bll.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)

	fmt.Println("lbs end")
}
