// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

// AppConfigs contains the core application configuration settings.
// It defines essential parameters that govern the application's behavior
// such as environment, logging settings, and secret management options.
type AppConfigs struct {
	//ENV: GO_ENV
	//
	// GoEnv specifies the environment in which the application runs (e.g., development, production)
	Environment Environment `mapstructure:"GO_ENV" envconfig:"GO_ENV"`
	//ENV: APP_NAME
	//
	//Name holds the name of the application for identification in logs and monitoring
	Name string `mapstructure:"APP_NAME" envconfig:"APP_NAME"`
	//ENV: APP_NAME
	//
	//Namespace is used to categorize application resources and configurations
	Namespace string `mapstructure:"NAMESPACE" envconfig:"NAMESPACE"`
	//ENV: NAMESPACE
	//
	//LogLevel determines the verbosity and severity threshold of application logs
	LogLevel LogLevel `mapstructure:"LOG_LEVEL" envconfig:"LOG_LEVEL"`
	//ENV: LOG_LEVEL
	//
	//UseSecretManager indicates whether the application should retrieve secrets from a secure secret manager
	SecretManagerKind SecretManagerKind `mapstructure:"SECRET_MANAGER_KIND" envconfig:"SECRET_MANAGER_KIND"`
	//ENV: APP_SECRET_KEY
	//
	//SecretKey identifies the key/path used to fetch secrets from the secret manager
	SecretKey string `mapstructure:"APP_SECRET_KEY" envconfig:"APP_SECRET_KEY"`
	//ENV: APP_HOST
	//
	//Host specifies the hostname or IP address on which the application listens for incoming requests
	Host string `mapstructure:"APP_HOST" envconfig:"APP_HOST"`
	//ENV: APP_PORT
	//
	//Port defines the network port on which the application listens for incoming requests
	Port int `mapstructure:"APP_PORT" envconfig:"APP_PORT"`
}

const (
	AppGoEnvKey                = "GO_ENV"              // Application environment (development, production, etc.)
	AppNameEnvKey              = "APP_NAME"            // Application name
	AppNamespaceEnvKey         = "NAMESPACE"           // Namespace for application resources
	AppLogLevelEnvKey          = "LOG_LEVEL"           // Logging level (debug, info, warn, error, etc.)
	AppSecretManagerKindEnvKey = "SECRET_MANAGER_KIND" // Secret manager type (e.g., AWS Secrets Manager, HashiCorp Vault)
	AppSecretKeyEnvKey         = "APP_SECRET_KEY"      // Secret key for fetching secrets
	AppHostEnvKey              = "APP_HOST"            // Hostname or IP address for the application
	AppPortEnvKey              = "APP_PORT"            // Port for the application
)
