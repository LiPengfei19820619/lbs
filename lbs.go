package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/nakagami/firebirdsql"

	"zte/ims/lbs/bll/SecurityCtrl"
)

func main() {
	fmt.Println("lbs start running ... ")

	go SecurityCtrl.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)

	fmt.Println("lbs end")
}
