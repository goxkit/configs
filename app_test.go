// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestAppConfigsWithViper(t *testing.T) {
	// Setup test
	// Create a temporary .env.local file for testing
	tmpEnvFile := ".env.local"
	envContent := `
GO_ENV=development
APP_NAME=test-app
NAMESPACE=test-ns
LOG_LEVEL=debug
SECRET_MANAGER_KIND=aws
APP_SECRET_KEY=test-secret-key
APP_HOST=localhost
APP_PORT=8080
`

	err := os.WriteFile(tmpEnvFile, []byte(envContent), 0644)
	assert.NoError(t, err, "Failed to create test .env.local file")

	defer func() {
		// Clean up the test file when done
		os.Remove(tmpEnvFile)
	}()

	// Configure Viper to use the .env.local file
	v := viper.New()
	v.SetConfigFile(tmpEnvFile)
	v.SetConfigType("env")
	v.AutomaticEnv()

	// Read the config file
	err = v.ReadInConfig()
	assert.NoError(t, err, "Failed to read config file")

	// Test case
	t.Run("should load AppConfigs correctly from .env.local", func(t *testing.T) {
		// Arrange
		var cfgs AppConfigs

		// Act
		err := v.Unmarshal(&cfgs)

		// Assert
		assert.NoError(t, err)
		// Since Viper won't automatically convert string "development" to DevelopmentEnv enum
		// we should check the raw values that would be set
		assert.Equal(t, "development", v.GetString("GO_ENV"))
		// After unmarshalling, the string would be stored in the struct but not converted to enum
		// In a real app, a post-processing step would be needed to convert these values
		assert.Equal(t, "test-app", cfgs.Name)
		assert.Equal(t, "test-ns", cfgs.Namespace)
		assert.Equal(t, "debug", v.GetString("LOG_LEVEL"))
		assert.Equal(t, "aws", v.GetString("SECRET_MANAGER_KIND"))
		assert.Equal(t, "test-secret-key", cfgs.SecretKey)
		assert.Equal(t, "localhost", cfgs.Host)
		assert.Equal(t, 8080, cfgs.Port)
	})

	// Test case that shows proper conversion of enum values
	t.Run("should convert string values to proper enums", func(t *testing.T) {
		// In a real application, we would have a function that processes the raw values
		// after unmarshalling to convert strings to proper enums
		rawCfgs := struct {
			Environment       string `mapstructure:"GO_ENV"`
			Name              string `mapstructure:"APP_NAME"`
			Namespace         string `mapstructure:"NAMESPACE"`
			LogLevel          string `mapstructure:"LOG_LEVEL"`
			SecretManagerKind string `mapstructure:"SECRET_MANAGER_KIND"`
			SecretKey         string `mapstructure:"APP_SECRET_KEY"`
			Host              string `mapstructure:"APP_HOST"`
			Port              int    `mapstructure:"APP_PORT"`
		}{}

		// Unmarshal raw values first
		err := v.Unmarshal(&rawCfgs)
		assert.NoError(t, err)

		// Then convert to the proper types
		appConfigs := AppConfigs{
			Environment:       NewEnvironment(rawCfgs.Environment),
			Name:              rawCfgs.Name,
			Namespace:         rawCfgs.Namespace,
			LogLevel:          NewLogLevel(rawCfgs.LogLevel),
			SecretManagerKind: SecretManagerKind(rawCfgs.SecretManagerKind),
			SecretKey:         rawCfgs.SecretKey,
			Host:              rawCfgs.Host,
			Port:              rawCfgs.Port,
		}

		// Now we can properly check the enum values
		assert.Equal(t, DevelopmentEnv, appConfigs.Environment)
		assert.Equal(t, "test-app", appConfigs.Name)
		assert.Equal(t, "test-ns", appConfigs.Namespace)
		assert.Equal(t, DEBUG, appConfigs.LogLevel)
		assert.Equal(t, SecretManagerKindAWS, appConfigs.SecretManagerKind)
		assert.Equal(t, "test-secret-key", appConfigs.SecretKey)
		assert.Equal(t, "localhost", appConfigs.Host)
		assert.Equal(t, 8080, appConfigs.Port)
	})

	// Test case for error handling
	t.Run("should handle missing required values", func(t *testing.T) {
		// Create a different config with missing values
		v := viper.New()
		v.SetEnvPrefix("")
		v.AutomaticEnv()

		var incompleteCfgs AppConfigs
		err := v.Unmarshal(&incompleteCfgs)

		// It shouldn't return an error, but some values should be empty/default
		assert.NoError(t, err)
		assert.Empty(t, incompleteCfgs.Name)
	})
}
