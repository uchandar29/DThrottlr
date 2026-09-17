package main

import (
	"log"

	"github.com/uchandar29/DThrottlr/internal/config"
	"github.com/uchandar29/DThrottlr/internal/server"
)

func main() {
	log.Println("Starting Dthrottlr server...")

	/* Load Configuration */
	appConf, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		return
	}
	log.Printf("Configuration loaded: %+v", appConf)

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
