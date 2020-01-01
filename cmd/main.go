package main

import (
	"fmt"
	"log"
	"os"
	"rebu/pkg/common"
	"strings"
)

func main() {
	listenAddress := strings.Trim(os.Getenv("APP_LISTENER_ADDR"), " ")
	if listenAddress == "" {
		log.Fatal("environment variable APP_LISTENER_ADDR not set")
	}

	err := verifyEnvironment()
	if err != nil {
		log.Fatalf("Error occurred while checking env variables: %v", err)
	}

	common.Start(listenAddress)
	fmt.Println("Server listening on: ", listenAddress)
	select {}
}

func verifyEnvironment() error {
	if strings.Trim(os.Getenv("DB_USERNAME"), " ") == "" {
		return fmt.Errorf("Env variable DB_USERNAME was not set")
	}

	if strings.Trim(os.Getenv("DB_PASSWORD"), " ") == "" {
		return fmt.Errorf("Env variable DB_PASSWORD was not set")
	}

	if strings.Trim(os.Getenv("REDIS_LISTENER_ADDR"), " ") == "" {
		return fmt.Errorf("Env variable REDIS_LISTENER_ADDR was not set")
	}

	return nil
}
