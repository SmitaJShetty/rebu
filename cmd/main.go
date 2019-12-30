package main

import (
	"fmt"
	"../pkg/common"
	"os"
)

func main() {
	listenAddress := os.Getenv("APP_LISTENER_PORT") //"localhost:8090"
	common.Start(listenAddress)
	fmt.Println("Server listening on: ", listenAddress)
	select {}
}
