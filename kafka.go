// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

// KafkaConfigs defines configuration parameters for Apache Kafka connections.
// It provides settings for connection, authentication, and security protocols.
type KafkaConfigs struct {
	//ENV: KAFKA_HOST
	//
	// Host specifies the Kafka broker hostname or IP address
	Host string `mapstructure:"KAFKA_HOST" envconfig:"KAFKA_HOST" default:"localhost"`
	//ENV: KAFKA_PORT
	//
	// Port defines the network port on which the Kafka broker is listening
	Port int `mapstructure:"KAFKA_PORT" envconfig:"KAFKA_PORT" default:"9092"`
	//ENV: KAFKA_TIMEOUT
	//
	// Timeout specifies the maximum time to wait for a response from the Kafka broker
	Timeout int `mapstructure:"KAFKA_TIMEOUT" envconfig:"KAFKA_TIMEOUT" default:"30s"`
	//ENV: KAFKA_SECURITY_PROTOCOL
	//
	// SecurityProtocol defines the protocol used to communicate with brokers
	// (e.g., "PLAINTEXT", "SSL", "SASL_PLAINTEXT", "SASL_SSL")
	SecurityProtocol string `mapstructure:"KAFKA_SECURITY_PROTOCOL" envconfig:"KAFKA_SECURITY_PROTOCOL" default:"SASL_PLAINTEXT"`
	//ENV: KAFKA_SASL_MECHANISMS
	//
	// SASLMechanisms specifies the SASL mechanism to use for authentication
	// (e.g., "PLAIN", "SCRAM-SHA-256", "SCRAM-SHA-512")
	SASLMechanisms string `mapstructure:"KAFKA_SASL_MECHANISMS" envconfig:"KAFKA_SASL_MECHANISMS" default:"PLAIN"`
	//ENV: KAFKA_CERTIFICATE_PATH
	//
	// CertificatePath specifies the path to the SSL certificate file for secure connections
	CertificatePath string `mapstructure:"KAFKA_CERTIFICATE_PATH" envconfig:"KAFKA_CERTIFICATE_PATH"`
	//ENV: KAFKA_TRUSTSTORE_PATH
	//
	// TrustStorePath specifies the path to the SSL truststore file for secure connections
	TrustStorePath string `mapstructure:"KAFKA_TRUSTSTORE_PATH" envconfig:"KAFKA_TRUSTSTORE_PATH"`
	//ENV: KAFKA_KEYSTORE_PATH
	//
	// KeyStorePath specifies the path to the SSL keystore file for secure connections
	KeyStorePath string `mapstructure:"KAFKA_KEYSTORE_PATH" envconfig:"KAFKA_KEYSTORE_PATH"`
	//ENV: KAFKA_ENDPOINT_IDENTIFICATION_ALGORITHM
	//
	// EndpointIdentificationAlgorithm specifies the algorithm used for hostname verification
	EndpointIdentificationAlgorithm string `mapstructure:"KAFKA_ENDPOINT_IDENTIFICATION_ALGORITHM" envconfig:"KAFKA_ENDPOINT_IDENTIFICATION_ALGORITHM"`
	//ENV: KAFKA_USER
	//
	// User specifies the username for Kafka broker authentication when using SASL
	User string `mapstructure:"KAFKA_USER" envconfig:"KAFKA_USER" default:"guest"`
	//ENV: KAFKA_PASSWORD
	//
	// Password contains the authentication credential for the Kafka user
	Password string `mapstructure:"KAFKA_PASSWORD" envconfig:"KAFKA_PASSWORD" default:"guest"`
}
