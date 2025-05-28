// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

import "time"

type OTLPConfigs struct {
	//ENV: OTLP_ENABLED
	//
	// Enabled indicates whether OTLP exporter is enabled, default is false
	Enabled bool `mapstructure:"OTLP_ENABLED" envconfig:"OTLP_ENABLED"`

	//ENV: OTLP_EXPORTER_TYPE
	//
	// ExporterType specifies the type of OTLP exporter to use (e.g., "otlp", "otlphttp")
	ExporterType string `mapstructure:"OTLP_EXPORTER_TYPE" envconfig:"OTLP_EXPORTER_TYPE"`
	//ENV: OTLP_ENDPOINT
	//
	// Endpoint specifies the endpoint URL for the OTLP exporter
	Endpoint string `mapstructure:"OTLP_ENDPOINT" envconfig:"OTLP_ENDPOINT" default:"localhost:4317"`
	//ENV: OTLP_ACCESS_KEY
	//
	// AccessKey is the access key for authentication with the OTLP endpoint
	AccessKey string `mapstructure:"OTLP_ACCESS_KEY" envconfig:"OTLP_ACCESS_KEY"`
	//ENV: OTLP_EXPORTER_TIMEOUT
	//
	// ExporterTimeout specifies the timeout duration for the OTLP exporter
	ExporterTimeout time.Duration `mapstructure:"OTLP_EXPORTER_TIMEOUT" envconfig:"OTLP_EXPORTER_TIMEOUT" default:"10s"`
	//ENV: OTLP_EXPORTER_INTERVAL
	//
	// ExporterInterval specifies the interval duration for the OTLP exporter
	// Default: 10s
	ExporterInterval time.Duration `mapstructure:"OTLP_EXPORTER_INTERVAL" envconfig:"OTLP_EXPORTER_INTERVAL" default:"30s"`
	//ENV: OTLP_EXPORTER_RECONNECTION_PERIOD
	//
	// ExporterReconnectionPeriod specifies the period for reconnection attempts
	ExporterReconnectionPeriod time.Duration `mapstructure:"OTLP_EXPORTER_RECONNECTION_PERIOD" envconfig:"OTLP_EXPORTER_RECONNECTION_PERIOD" default:"30s"`
	//ENV: OTLP_METRICS_EXPORTER_RATE_BASE
	//
	// MetricsExporterRateBase specifies the rate base for metrics exporter
	MetricsExporterRateBase float64 `mapstructure:"OTLP_METRICS_EXPORTER_RATE_BASE" envconfig:"OTLP_METRICS_EXPORTER_RATE_BASE" default:"0.8"`
	//ENV: OTLP_TRACING_EXPORTER_RATE_BASE
	//
	// TracingExporterRateBase specifies the rate base for tracing exporter
	TracingExporterRateBase float64 `mapstructure:"OTLP_TRACING_EXPORTER_RATE_BASE" envconfig:"OTLP_TRACING_EXPORTER_RATE_BASE" default:"0.8"`
	// ENV: OTLP_METRICS_ENABLED
	//
	// MetricsEnabled indicates whether metrics collection is enabled
	MetricsEnabled bool `mapstructure:"OTLP_METRICS_ENABLED" envconfig:"OTLP_METRICS_ENABLED"`
	// ENV: OTLP_TRACING_ENABLED
	//
	// TracingEnabled indicates whether tracing collection is enabled
	TracingEnabled bool `mapstructure:"OTLP_TRACING_ENABLED" envconfig:"OTLP_TRACING_ENABLED"`
}
