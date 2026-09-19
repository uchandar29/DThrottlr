package main

import (
	"fmt"
	"log"

	"github.com/uchandar29/DThrottlr/internal/config"
	"github.com/uchandar29/DThrottlr/internal/limiter"
	"github.com/uchandar29/DThrottlr/internal/redisclient"
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

	/* Setup Redis Connection */
	rdsClient, err := redisclient.New(appConf.RedisConfig)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		return
	}
	defer rdsClient.Close()

	limiter := limiter.New(rdsClient, appConf.LimiterConfig.BucketSize, appConf.LimiterConfig.TokenRefillRate)

	if err := server.Run(fmt.Sprintf(":%d", appConf.ServerConfig.Port), limiter); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
