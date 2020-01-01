package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Config holds service library config
type Config struct {
	// Kafka
	KafkaHost                  string
	KafkaPort                  string
	KafkaWebsocketPort         string
	KafkaProducerPool          PoolConfig
	KafkaWatermarkConsumerPool PoolConfig
	// Redis
	RedisHost           string
	RedisPort           string
	RedisConnectionPool PoolConfig
	// RPC
	RPCHost           string
	RPCPort           string
	RPCTimeoutSeconds float32
	RPCConnectionPool PoolConfig
	// Heartbeats
	HeartbeatInterval    time.Duration
	NoHeartbeatDeathTime time.Duration
}

// PoolConfig holds configuration for a resource pool
type PoolConfig struct {
	InitialCapacity    int
	MaxCapacity        int
	IdleTimeout        time.Duration
	PrefillParallelism int
}

func (p PoolConfig) String() string {
	return fmt.Sprintf("(Capacity: [%v, %v], IdleTimeout: %v, PrefilParallelism: %v)",
		p.InitialCapacity, p.MaxCapacity, p.IdleTimeout, p.PrefillParallelism)
}

func init() {
	defaultConfig = Config{
		// Kafka
		KafkaHost:          os.Getenv("KAFKA_HOST"),
		KafkaPort:          os.Getenv("KAFKA_PORT"),
		KafkaWebsocketPort: os.Getenv("KAFKA_WEBSOCKET_PORT"),
		KafkaProducerPool: PoolConfig{
			InitialCapacity:    10,
			MaxCapacity:        100,
			IdleTimeout:        1 * time.Minute,
			PrefillParallelism: 0,
		},
		KafkaWatermarkConsumerPool: PoolConfig{
			InitialCapacity:    10,
			MaxCapacity:        100,
			IdleTimeout:        1 * time.Minute,
			PrefillParallelism: 0,
		},
		// Redis
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		RedisConnectionPool: PoolConfig{
			InitialCapacity:    10,
			MaxCapacity:        100,
			IdleTimeout:        1 * time.Minute,
			PrefillParallelism: 0,
		},
		// RPC
		RPCHost:           os.Getenv("RPC_HOST"),
		RPCPort:           os.Getenv("RPC_PORT"),
		RPCTimeoutSeconds: 20.0,
		RPCConnectionPool: PoolConfig{
			InitialCapacity:    10,
			MaxCapacity:        100,
			IdleTimeout:        1 * time.Minute,
			PrefillParallelism: 0,
		},
		// Heartbeats
		HeartbeatInterval:    5 * time.Second,
		NoHeartbeatDeathTime: 15 * time.Second,
	}
}

// ToString converts a config to string for logging purposes
func (c *Config) ToString() string {
	r := []string{}
	r = append(r, fmt.Sprint(" ---------------------------------------------------------------------------"))
	r = append(r, fmt.Sprint(" | Service Library configuration"))
	r = append(r, fmt.Sprint(" ------------ Kafka ---------------------------------------------------------"))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "KafkaHost", c.KafkaHost))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "KafkaPort", c.KafkaPort))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "KafkaWebsocketPort", c.KafkaWebsocketPort))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "KafkaProducerPool", c.KafkaProducerPool))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "KafkaWatermarkConsumerPool", c.KafkaWatermarkConsumerPool))
	r = append(r, fmt.Sprint(" ------------ Redis ---------------------------------------------------------"))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "RedisHost", c.RedisHost))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "RedisPort", c.RedisPort))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "RedisConnectionPool", c.RedisConnectionPool))
	r = append(r, fmt.Sprint(" ------------- RPC ----------------------------------------------------------"))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "RPCHost", c.RPCHost))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "RPCPort", c.RPCPort))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "RPCTimeoutSeconds", c.RPCTimeoutSeconds))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "RPCConnectionPool", c.RPCConnectionPool))
	r = append(r, fmt.Sprint(" ------------- Heartbeats ---------------------------------------------------"))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "HeartbeatInterval", c.HeartbeatInterval))
	r = append(r, fmt.Sprintf(" | %-30v | %v", "NoHeartbeatDeathTime", c.NoHeartbeatDeathTime))
	r = append(r, fmt.Sprint("\n"))
	return strings.Join(r, "\n")
}

// Default config to be initialized on startup and used every time a new config is created
var defaultConfig Config

// New creates a new Config by providing default values for any values not provided in CustomConfig.
func New() Config {
	c := defaultConfig
	return c
}
