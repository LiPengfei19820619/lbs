package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/nakagami/firebirdsql"

	"zte/ims/lbs/bll/SecurityCtrl"
	"zte/ims/lbs/log"
)

func main() {
	fmt.Println("lbs start running ... ")
	log.WriteRunLog("aaa")

	go SecurityCtrl.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)

	fmt.Println("lbs end")
}
