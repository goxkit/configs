// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

// Package configs provides a comprehensive configuration framework for GoKit applications.
// It offers a structured approach to manage different types of configurations like HTTP,
// database connections, message brokers, authentication, and more.
package configs

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Configs is the central configuration container that aggregates all available
// configuration components in the GoKit framework. It allows applications to
// maintain a single reference to all necessary configurations.
type Configs struct {
	// Logger provides access to the application's configured logger
	Logger *zap.Logger
	// Viper is the configuration management library used to load and manage
	Custom *viper.Viper

	// AppConfigs contains basic application configurations like environment and name
	AppConfigs *AppConfigs
	// HTTPConfigs provides configuration for HTTP servers/clients
	HTTPConfigs *HTTPConfigs
	//
	OTLPConfigs *OTLPConfigs
	// PostgresConfigs contains database connection and pool settings
	PostgresConfigs *PostgresConfigs
	// IdentityConfigs provides authentication and identity management settings
	IdentityConfigs *IdentityConfigs
	// MQTTConfigs holds MQTT message broker settings
	MQTTConfigs *MQTTConfigs
	// RabbitMQConfigs provides RabbitMQ message broker configuration
	RabbitMQConfigs *RabbitMQConfigs
	// KafkaConfigs contains Kafka message broker settings
	KafkaConfigs *KafkaConfigs
	// AWSConfigs holds AWS services configuration
	AWSConfigs *AWSConfigs
	// DynamoDBConfigs provides Amazon DynamoDB specific settings
	DynamoDBConfigs *DynamoDBConfigs
}
