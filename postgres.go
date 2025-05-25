// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

// SQLConfigs defines configuration parameters for SQL database connections.
// It contains all the necessary information to establish and maintain
// connections to SQL databases like PostgreSQL
type PostgresConfigs struct {
	//ENV: POSTGRES_HOST
	//
	// Host specifies the database server hostname or IP address
	Host string `mapstructure:"POSTGRES_HOST" envconfig:"POSTGRES_HOST"`
	//ENV: POSTGRES_PORT
	//
	// Port defines the network port on which the database server is listening
	Port string `mapstructure:"POSTGRES_PORT" envconfig:"POSTGRES_PORT"`
	//ENV: POSTGRES_USER
	//
	// User specifies the username for database authentication
	User string `mapstructure:"POSTGRES_USER" envconfig:"POSTGRES_USER"`
	//ENV: POSTGRES_PASSWORD
	//
	// Password contains the authentication credential for the database user
	Password string `mapstructure:"POSTGRES_PASSWORD" envconfig:"POSTGRES_PASSWORD"`
	//ENV: POSTGRES_DB_NAME
	//
	// DbName specifies the name of the database to connect to
	DbName string `mapstructure:"POSTGRES_DB_NAME" envconfig:"POSTGRES_DB_NAME"`
	//ENV: POSTGRES_SSL_ENABLED
	//
	// SslMode indicates whether to use SSL for the database connection
	SslEnabled bool `mapstructure:"POSTGRES_SSL_ENABLED" envconfig:"POSTGRES_SSL_ENABLED"`
	//ENV: POSTGRES_SECONDS_TO_PING
	//
	// SecondsToPing defines the interval in seconds between health check pings
	// to verify the database connection remains active
	SecondsToPing int `mapstructure:"POSTGRES_SECONDS_TO_PING" envconfig:"POSTGRES_SECONDS_TO_PING"`
}
