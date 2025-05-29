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
	//ENV: OTLP_ENDPOINT
	//
	// Endpoint specifies the endpoint URL for the OTLP exporter
	Endpoint string `mapstructure:"OTLP_ENDPOINT" envconfig:"OTLP_ENDPOINT" default:"localhost:4317"`
	//ENV: OTLP_EXPORTER_HEADERS
	//
	//ExporterHeaders specifies the headers to be sent with the OTLP exporter
	//
	// Example:
	//
	// OTLP_EXPORTER_HEADERS=api-key=key,other-config-value=value"
	ExporterHeaders string `mapstructure:"OTLP_EXPORTER_HEADERS" envconfig:"OTLP_EXPORTER_HEADERS"`
	//ENV: OTLP_EXPORTER_TLS_ENABLED
	//
	// ExporterTLSEnabled indicates whether to use SSL for the OTLP exporter
	ExporterTLSEnabled bool `mapstructure:"OTLP_EXPORTER_TLS_ENABLED" envconfig:"OTLP_EXPORTER_TLS_ENABLED"`
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
	//ENV: OTLP_EXPORTER_IDLE_TIMEOUT
	//
	// ExporterIdleTimeout specifies the idle timeout for the OTLP exporter
	ExporterIdleTimeout time.Duration `mapstructure:"OTLP_EXPORTER_IDLE_TIMEOUT" envconfig:"OTLP_EXPORTER_IDLE_TIMEOUT" default:"30s"`
	//ENV: OTLP_EXPORTER_KEEP_ALIVE_TIME
	//
	// ExporterKeepAlivePeriod specifies the keep-alive period for the OTLP exporter
	ExporterKeepAliveTime time.Duration `mapstructure:"OTLP_EXPORTER_KEEP_ALIVE_TIME" envconfig:"OTLP_EXPORTER_KEEP_ALIVE_TIME" default:"30s"`
	//ENV: OTLP_EXPORTER_KEEP_ALIVE_TIMEOUT
	//
	// ExporterKeepAliveTimeout specifies the keep-alive timeout for the OTLP exporter
	ExporterKeepAliveTimeout time.Duration `mapstructure:"OTLP_EXPORTER_KEEP_ALIVE_TIMEOUT" envconfig:"OTLP_EXPORTER_KEEP_ALIVE_TIMEOUT" default:"10s"`
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
