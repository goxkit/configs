// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

type SecretManagerKind string

const (
	// SecretManagerKindNone indicates that no secret manager is used
	SecretManagerKindNone SecretManagerKind = "none"
	// SecretManagerKindAWS indicates that AWS Secrets Manager is used
	SecretManagerKindAWS SecretManagerKind = "aws"
	//
	SecretManagerKindVault SecretManagerKind = "vault"
)
