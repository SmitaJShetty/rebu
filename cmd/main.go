package main

import (
	"fmt"
	"os"
	"rebu/pkg/common"
)

func main() {
	listenAddress := os.Getenv("APP_LISTENER_PORT")
	common.Start(listenAddress)
	fmt.Println("Server listening on: ", listenAddress)
	select {}
}
