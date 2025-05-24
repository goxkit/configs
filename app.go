// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

// AppConfigs contains the core application configuration settings.
// It defines essential parameters that govern the application's behavior
// such as environment, logging settings, and secret management options.
type AppConfigs struct {
	// GoEnv specifies the environment in which the application runs (e.g., development, production)
	Env Environment
	// Name holds the name of the application for identification in logs and monitoring
	Name string
	//
	Namespace string
	// LogLevel determines the verbosity and severity threshold of application logs
	LogLevel LogLevel
	// UseSecretManager indicates whether the application should retrieve secrets from a secure secret manager
	SecretManagerKind SecretManagerKind
	// SecretKey identifies the key/path used to fetch secrets from the secret manager
	SecretKey string
	//
	Host string
	//
	Port int
}
