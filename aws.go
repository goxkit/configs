// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

// AWSConfigs defines authentication and credential settings for AWS services.
// It contains the necessary parameters to authenticate with AWS APIs.
type AWSConfigs struct {
	//ENV: AWS_REGION
	//
	// Region specifies the AWS region where the service is located
	Region string `mapstructure:"AWS_REGION" envconfig:"AWS_REGION"`
	//ENV: AWS_ACCESS_KEY_ID
	//
	// AccessKeyID is the AWS access key part of the credential pair
	AccessKeyID string `mapstructure:"AWS_ACCESS_KEY_ID" envconfig:"AWS_ACCESS_KEY_ID"`
	//ENV: AWS_SECRET_ACCESS_KEY
	//
	// SecretAccessKey is the AWS secret key part of the credential pair
	SecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY" envconfig:"AWS_SECRET_ACCESS_KEY"`
	//ENV: AWS_SESSION_TOKEN
	//
	// SessionToken provides temporary credentials when using AWS STS (Security Token Service)
	SessionToken string `mapstructure:"AWS_SESSION_TOKEN" envconfig:"AWS_SESSION_TOKEN"`
}

// DynamoDBConfigs provides configuration settings specific to Amazon DynamoDB.
// It contains connection and targeting parameters for DynamoDB operations.
type DynamoDBConfigs struct {
	//ENV: DYNAMODB_ENDPOINT
	//
	// Endpoint specifies the DynamoDB service endpoint URL (useful for local development)
	Endpoint string `mapstructure:"DYNAMODB_ENDPOINT" envconfig:"DYNAMODB_ENDPOINT"`
	//ENV: DYNAMODB_REGION
	//
	// Region defines the AWS region where the DynamoDB table is located
	Region string `mapstructure:"DYNAMODB_REGION" envconfig:"DYNAMODB_REGION"`
	//ENV: DYNAMODB_TABLE
	//
	// Table specifies the default DynamoDB table name to use
	Table string `mapstructure:"DYNAMODB_TABLE" envconfig:"DYNAMODB_TABLE"`
}
