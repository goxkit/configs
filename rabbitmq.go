// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

// RabbitMQConfigs defines configuration parameters for RabbitMQ message broker connections.
// It contains all necessary connection details for establishing a connection to a RabbitMQ server.
type RabbitMQConfigs struct {
	//ENV: RABBITMQ_SCHEMA
	//
	// Schema defines the connection protocol (typically "amqp" or "amqps" for TLS)
	Schema string `mapstructure:"RABBITMQ_SCHEMA" envconfig:"RABBITMQ_SCHEMA" default:"amqp"`
	//ENV: RABBITMQ_HOST
	//
	// Host specifies the RabbitMQ server hostname or IP address
	Host string `mapstructure:"RABBITMQ_HOST" envconfig:"RABBITMQ_HOST" default:"localhost"`
	//ENV: RABBITMQ_PORT
	//
	// Port defines the network port on which the RabbitMQ server is listening
	Port string `mapstructure:"RABBITMQ_PORT" envconfig:"RABBITMQ_PORT" default:"5672"`
	//ENV: RABBITMQ_USERNAME
	//
	// User specifies the username for RabbitMQ server authentication
	User string `mapstructure:"RABBITMQ_USERNAME" envconfig:"RABBITMQ_USERNAME" default:"guest"`
	//ENV: RABBITMQ_PASSWORD
	//
	// Password contains the authentication credential for the RabbitMQ user
	Password string `mapstructure:"RABBITMQ_PASSWORD" envconfig:"RABBITMQ_PASSWORD" default:"guest"`
	//ENV: RABBITMQ_VHOST
	//
	// VHost specifies the virtual host to use on the RabbitMQ server
	VHost string `mapstructure:"RABBITMQ_VHOST" envconfig:"RABBITMQ_VHOST" default:"/"`
}

const (
	RabbitSchemaEnvKey   = "RABBIT_SCHEMA"   // RabbitMQ connection schema (amqp, amqps)
	RabbitHostEnvKey     = "RABBIT_HOST"     // RabbitMQ host
	RabbitPortEnvKey     = "RABBIT_PORT"     // RabbitMQ port
	RabbitUserEnvKey     = "RABBIT_USER"     // RabbitMQ username
	RabbitPasswordEnvKey = "RABBIT_PASSWORD" // RabbitMQ password
	RabbitVHostEnvKey    = "RABBIT_VHOST"    // RabbitMQ virtual host
)
