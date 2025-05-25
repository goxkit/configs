// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

// MQTTConfigs defines configuration parameters for MQTT broker connections.
// It provides settings for connection, authentication, and TLS/SSL security.
type MQTTConfigs struct {
	//ENV: MQTT_HOST
	//
	// Host specifies the MQTT broker hostname or IP address
	Host string `mapstructure:"MQTT_HOST" envconfig:"MQTT_HOST"`
	//ENV: MQTT_PORT
	//
	// Port defines the network port on which the MQTT broker is listening
	Port int `mapstructure:"MQTT_PORT" envconfig:"MQTT_PORT"`
	//ENV: MQTT_USER
	//
	// User specifies the username for MQTT broker authentication
	User string `mapstructure:"MQTT_USER" envconfig:"MQTT_USER"`
	//ENV: MQTT_PASSWORD
	//
	// Password contains the authentication credential for the MQTT user
	Password string `mapstructure:"MQTT_PASSWORD" envconfig:"MQTT_PASSWORD"`
	//ENV: MQTT_PROTOCOL
	//
	// Protocol specifies the MQTT protocol version to use (e.g., "mqtt", "mqtts")
	Protocol string `mapstructure:"MQTT_PROTOCOL" envconfig:"MQTT_PROTOCOL"`
	//ENV: MQTT_ROOT_CA_PATH
	//
	// RootCaPath is the file path to the root CA certificate for TLS verification
	RootCaPath string `mapstructure:"MQTT_ROOT_CA_PATH" envconfig:"MQTT_ROOT_CA_PATH"`
	//ENV: MQTT_CERT_PATH
	//
	// CertPath is the file path to the client certificate for mutual TLS authentication
	CertPath string `mapstructure:"MQTT_CERT_PATH" envconfig:"MQTT_CERT_PATH"`
	//ENV: MQTT_PRIVATE_KEY_PATH
	//
	// PrivateKeyPath is the file path to the private key associated with the client certificate
	PrivateKeyPath string `mapstructure:"MQTT_PRIVATE_KEY_PATH" envconfig:"MQTT_PRIVATE_KEY_PATH"`
}

const (
	MQTTProtocolEnvKey = "MQTT_PROTOCOL" // MQTT protocol (mqtt, mqtts)
	MQTTHostEnvKey     = "MQTT_HOST"     // MQTT broker host
	MQTTPortEnvKey     = "MQTT_PORT"     // MQTT broker port
	MQTTUserEnvKey     = "MQTT_USER"     // MQTT username
	MQTTPasswordEnvKey = "MQTT_PASSWORD" // MQTT password
)
