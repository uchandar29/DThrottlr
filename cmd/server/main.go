package main

import (
	"log"

	"github.com/uchandar29/DThrottlr/internal/server"
)

func main() {
	log.Println("Starting Dthrottlr server...")

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
