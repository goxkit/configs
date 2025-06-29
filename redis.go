// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

type RedisConfigs struct {
	//ENV: REDIS_HOST
	//
	// Host specifies the Redis server hostname or IP address
	Host string `mapstructure:"REDIS_HOST" envconfig:"REDIS_HOST" default:"localhost"`
	//ENV: REDIS_PORT
	//
	// Port defines the network port on which the Redis server is listening
	Port int `mapstructure:"REDIS_PORT" envconfig:"REDIS_PORT" default:"6379"`
	//ENV: REDIS_PASSWORD
	//
	// Password contains the authentication credential for the Redis server
	Password string `mapstructure:"REDIS_PASSWORD" envconfig:"REDIS_PASSWORD" default:"guest"`
	//ENV: REDIS_DB
	//
	// Db specifies the Redis database number to connect to
	Db int `mapstructure:"REDIS_DB" envconfig:"REDIS_DB" default:"0"`
	//ENV: REDIS_TLS_ENABLED
	//
	// TLSEnabled indicates whether to use TLS for the Redis connection
	TLSEnabled bool `mapstructure:"REDIS_TLS_ENABLED" envconfig:"REDIS_TLS_ENABLED"`
	//ENV: REDIS_CLUSTER_ENABLED
	//
	// ClusterEnabled indicates whether Redis cluster mode is enabled
	ClusterEnabled bool `mapstructure:"REDIS_CLUSTER_ENABLED" envconfig:"REDIS_CLUSTER_ENABLED"`
}
