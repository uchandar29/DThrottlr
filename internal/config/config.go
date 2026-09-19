package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type LimiterConfig struct {
	BucketSize      int `yaml:"bucket_size"`
	TokenRefillRate int `yaml:"token_refill_rate"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

type NodeConfig struct {
	NodeCount string `yaml:"node_count"`
	BasePort  string `yaml:"base_port"`
}

type Config struct {
	ServerConfig   ServerConfig   `yaml:"server"`
	RedisConfig    RedisConfig    `yaml:"redis"`
	LimiterConfig  LimiterConfig  `yaml:"limiter"`
	DatabaseConfig DatabaseConfig `yaml:"database"`
	NodesConfig    NodeConfig     `yaml:"nodes"`
}

func LoadConfig(path string) (*Config, error) {
	yamlConfigData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(yamlConfigData, &config); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal config: %w", err)
	}

	return &config, nil
}
