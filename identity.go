// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

// IdentityConfigs contains configuration settings for identity and authentication services.
// It provides the necessary parameters for OAuth2/OIDC flows and JWT token verification.
type IdentityConfigs struct {
	//ENV: IDENTITY_CLIENT_ID
	//
	// ClientID represents the application's client identifier for authentication with identity providers
	ClientID string `mapstructure:"IDENTITY_CLIENT_ID" envconfig:"IDENTITY_CLIENT_ID"`
	//ENV: IDENTITY_CLIENT_SECRET
	//
	// ClientSecret contains the secret key used to authenticate the application with the identity provider
	ClientSecret string `mapstructure:"IDENTITY_CLIENT_SECRET" envconfig:"IDENTITY_CLIENT_SECRET"`
	//ENV: IDENTITY_GRANT_TYPE
	//
	// GrantType specifies the OAuth2 grant type to be used (e.g., "client_credentials", "authorization_code")
	GrantType string `mapstructure:"IDENTITY_GRANT_TYPE" envconfig:"IDENTITY_GRANT_TYPE"`
	//ENV: IDENTITY_MILLISECONDS_BETWEEN_JWK
	//
	// MillisecondsBetweenJWK defines the interval between JWK (JSON Web Key) refresh operations
	MillisecondsBetweenJWK int64 `mapstructure:"IDENTITY_MILLISECONDS_BETWEEN_JWK" envconfig:"IDENTITY_MILLISECONDS_BETWEEN_JWK"`
	//ENV: IDENTITY_DOMAIN
	//
	// Domain specifies the identity provider's domain (often used with Auth0 or similar services)
	Domain string `mapstructure:"IDENTITY_DOMAIN" envconfig:"IDENTITY_DOMAIN"`
	//ENV: IDENTITY_AUDIENCE
	//
	// Audience identifies the intended recipient of the authentication tokens
	Audience string `mapstructure:"IDENTITY_AUDIENCE" envconfig:"IDENTITY_AUDIENCE"`
	//ENV: IDENTITY_ISSUER
	//
	// Issuer identifies the entity that issued the JWT token
	Issuer string `mapstructure:"IDENTITY_ISSUER" envconfig:"IDENTITY_ISSUER"`
	//ENV: IDENTITY_SIGNATURE
	//
	// Signature contains cryptographic material used to verify token signatures
	Signature string `mapstructure:"IDENTITY_SIGNATURE" envconfig:"IDENTITY_SIGNATURE"`
}
