package main

import (
	"fmt"
	"../pkg/common"
)

func main() {
	listenAddress := "localhost:8090"
	common.Start(listenAddress)
	fmt.Println("Server listening on: ", listenAddress)
	select {}
}
