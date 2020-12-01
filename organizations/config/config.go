package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Config represents application wide configuration parameters.
type Config struct {
	QueueGroupName string
	NATSURL        string
	NATSClientID   string
	NATSClusterID  string
	AckWait        time.Duration

	PGConnString   string
}

var envConfig *Config

// Init initializes the config object with the environmental variables present at the time of its calling.
func Init() *Config {
	return setConfig()
}

// GetConfig returns the configuration object. It is not guarenteed to be initialized.
func GetConfig() *Config {
	return envConfig
}

func get(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Please set the %s environmental variable.", key)
	}
	return value
}

func setConfig() *Config {
	envConfig = &Config{
		QueueGroupName: get("NATS_QUEUE_GROUP"),
		NATSURL: get("NATS_URL"),
		NATSClientID: get("NATS_CLIENT_ID"),
		NATSClusterID: get("NATS_CLUSTER_ID"),
		AckWait: time.Duration(5) * time.Second,

		PGConnString: getPGConnString(),
	}
	return envConfig
}

func getPGConnString() string {
	host := get("PG_HOST")
	user := get("PG_USER")
	password := get("PG_PASSWORD")
	database := get("PG_DATABASE")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", host, user, password, database)
}