package main

import (
	"context"
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

	// -- Test the limiter with a sample client ID
	lim := limiter.New(rdsClient, appConf.LimiterConfig.BucketSize, appConf.LimiterConfig.TokenRefillRate)
	allowed, remaining, _ := lim.Allow(context.Background(), "test-client")
	log.Printf("allowed=%v remaining=%d", allowed, remaining)
	// -- Test the limiter with a sample client ID

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
