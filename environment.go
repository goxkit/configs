// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

import "fmt"

// Environment represents the application execution environment as a typed enum.
// It provides type safety when dealing with different deployment environments.
type Environment string

const (
	// UnknownEnv represents an unspecified or invalid environment
	UnknownEnv Environment = "unknown"
	// LocalEnv represents a local development environment
	LocalEnv Environment = "local"
	// DevelopmentEnv represents a shared development environment
	DevelopmentEnv Environment = "development"
	// StagingEnv represents a pre-production environment for testing
	StagingEnv Environment = "staging"
	// QaEnv represents a quality assurance testing environment
	QaEnv Environment = "qa"
	// ProductionEnv represents the live production environment
	ProductionEnv Environment = "production"
)

var (
	// EnvironmentMapping maps Environment enum values to their string representations
	EnvironmentMapping = map[Environment]string{
		UnknownEnv:     "unknown",
		LocalEnv:       "local",
		DevelopmentEnv: "development",
		StagingEnv:     "staging",
		QaEnv:          "qa",
		ProductionEnv:  "production",
	}
)

// NewEnvironment converts a string environment name to the corresponding Environment enum value.
// It accepts various standard naming conventions like "dev"/"development", case-insensitive.
// Returns UnknownEnv if the string doesn't match any known environment.
func NewEnvironment(env string) Environment {
	switch env {
	case "local":
		fallthrough
	case "LOCAL":
		return LocalEnv
	case "development":
		fallthrough
	case "DEVELOPMENT":
		fallthrough
	case "dev":
		fallthrough
	case "DEV":
		return DevelopmentEnv
	case "production":
		fallthrough
	case "PRODUCTION":
		fallthrough
	case "prod":
		fallthrough
	case "PROD":
		return ProductionEnv
	case "staging":
		fallthrough
	case "STAGING":
		fallthrough
	case "stg":
		fallthrough
	case "STG":
		return StagingEnv
	case "qa":
		fallthrough
	case "QA":
		return QaEnv
	default:
		return UnknownEnv
	}
}

func (e Environment) EnvFile() string {
	switch e {
	case LocalEnv:
		return ".env.local"
	case DevelopmentEnv:
		return ".env.dev"
	case ProductionEnv:
		return ".env.prod"
	case StagingEnv:
		return ".env.stg"
	case QaEnv:
		return ".env.qa"
	default:
		return ".env"
	}
}

// ToString returns the canonical string representation of the Environment.
// This provides consistent environment naming when logging or presenting the environment.
func (e Environment) ToString() string {
	switch e {
	case LocalEnv:
		return "local"
	case DevelopmentEnv:
		return "development"
	case ProductionEnv:
		return "production"
	case StagingEnv:
		return "staging"
	case QaEnv:
		return "qa"
	default:
		return "unknown"
	}
}

func (e *Environment) String() string {
	return e.ToString()
}

func (e *Environment) Unmarshal(v interface{}) error {
	switch v := v.(type) {
	case string:
		*e = NewEnvironment(v)
		return nil
	default:
		return fmt.Errorf("unsupported type %T for Environment", v)
	}
}

func (e Environment) Marshal() (interface{}, error) {
	return e.ToString(), nil
}
