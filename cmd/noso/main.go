package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Noso app starting...")

	if err := initializeApp(); err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	startServer()
}

func initializeApp() error {
	return nil
}

func startServer() {
	fmt.Println("Server is running...")
}
