package config

import "time"

// Config represents application wide configuration parameters.
type Config struct {
	QueueGroupName string
	NATSURL        string
	AckWait        time.Duration
}