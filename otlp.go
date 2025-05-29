// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

import "time"

type OTLPConfigs struct {
	//ENV: OTEL_EXPORTER_OTLP_ENABLED
	//
	// Enabled indicates whether OTLP exporter is enabled, default is false
	Enabled bool `mapstructure:"OTEL_EXPORTER_OTLP_ENABLED" envconfig:"OTEL_EXPORTER_OTLP_ENABLED"`
	//ENV: OTEL_EXPORTER_OTLP_ENDPOINT
	//
	// Endpoint specifies the endpoint URL for the OTLP exporter
	Endpoint string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT" envconfig:"OTEL_EXPORTER_OTLP_ENDPOINT" default:"localhost:4317"`
	//ENV: OTEL_EXPORTER_OTLP_HEADERS
	//
	//ExporterHeaders specifies the headers to be sent with the OTLP exporter
	//
	// Example:
	//
	// OTEL_EXPORTER_OTLP_HEADERS=api-key=key,other-config-value=value"
	ExporterHeaders string `mapstructure:"OTEL_EXPORTER_OTLP_HEADERS" envconfig:"OTEL_EXPORTER_OTLP_HEADERS"`
	//ENV: OTEL_EXPORTER_OTLP_TLS_ENABLED
	//
	// ExporterTLSEnabled indicates whether to use SSL for the OTLP exporter
	ExporterTLSEnabled bool `mapstructure:"OTEL_EXPORTER_OTLP_TLS_ENABLED" envconfig:"OTEL_EXPORTER_OTLP_TLS_ENABLED"`
	//ENV: OTEL_EXPORTER_OTLP_TIMEOUT
	//
	// ExporterTimeout specifies the timeout duration for the OTLP exporter
	ExporterTimeout time.Duration `mapstructure:"OTEL_EXPORTER_OTLP_TIMEOUT" envconfig:"OTEL_EXPORTER_OTLP_TIMEOUT" default:"10s"`
	//ENV:  OTEL_EXPORTER_OTLP_INTERVAL
	//
	// ExporterInterval specifies the interval duration for the OTLP exporter
	// Default: 10s
	ExporterInterval time.Duration `mapstructure:"OTEL_EXPORTER_OTLP_INTERVAL" envconfig:"OTEL_EXPORTER_OTLP_INTERVAL" default:"30s"`
	//ENV: OTEL_EXPORTER_OTLP_RECONNECTION_PERIOD
	//
	// ExporterReconnectionPeriod specifies the period for reconnection attempts
	ExporterReconnectionPeriod time.Duration `mapstructure:"OTEL_EXPORTER_OTLP_RECONNECTION_PERIOD" envconfig:"OTEL_EXPORTER_OTLP_RECONNECTION_PERIOD" default:"30s"`
	//ENV: OTEL_EXPORTER_OTLP_IDLE_TIMEOUT
	//
	// ExporterIdleTimeout specifies the idle timeout for the OTLP exporter
	ExporterIdleTimeout time.Duration `mapstructure:"OTEL_EXPORTER_OTLP_IDLE_TIMEOUT" envconfig:"OTEL_EXPORTER_OTLP_IDLE_TIMEOUT" default:"30s"`
	//ENV: OTEL_EXPORTER_OTLP_KEEP_ALIVE_TIME
	//
	// ExporterKeepAlivePeriod specifies the keep-alive period for the OTLP exporter
	ExporterKeepAliveTime time.Duration `mapstructure:"OTEL_EXPORTER_OTLP_KEEP_ALIVE_TIME" envconfig:"OTEL_EXPORTER_OTLP_KEEP_ALIVE_TIME" default:"30s"`
	//ENV: OTEL_EXPORTER_OTLP_KEEP_ALIVE_TIMEOUT
	//
	// ExporterKeepAliveTimeout specifies the keep-alive timeout for the OTLP exporter
	ExporterKeepAliveTimeout time.Duration `mapstructure:"OTEL_EXPORTER_OTLP_KEEP_ALIVE_TIMEOUT" envconfig:"OTEL_EXPORTER_OTLP_KEEP_ALIVE_TIMEOUT" default:"10s"`
	//ENV: OTEL_EXPORTER_OTLP_METRICS_RATE_BASE
	//
	// MetricsExporterRateBase specifies the rate base for metrics exporter
	MetricsExporterRateBase float64 `mapstructure:"OTEL_EXPORTER_OTLP_METRICS_RATE_BASE" envconfig:"OTEL_EXPORTER_OTLP_METRICS_RATE_BASE" default:"0.8"`
	//ENV: OTEL_EXPORTER_OTLP_TRACES_RATE_BASE
	//
	// TracingExporterRateBase specifies the rate base for tracing exporter
	TracingExporterRateBase float64 `mapstructure:"OTEL_EXPORTER_OTLP_TRACES_RATE_BASE" envconfig:"OTEL_EXPORTER_OTLP_TRACES_RATE_BASE" default:"0.8"`
	// ENV: OTEL_EXPORTER_OTLP_METRICS_ENABLED
	//
	// MetricsEnabled indicates whether metrics collection is enabled
	MetricsEnabled bool `mapstructure:"OTEL_EXPORTER_OTLP_METRICS_ENABLED" envconfig:"OTEL_EXPORTER_OTLP_METRICS_ENABLED"`
	// ENV: OTEL_EXPORTER_OTLP_ENABLED
	//
	// TracingEnabled indicates whether tracing collection is enabled
	TracingEnabled bool `mapstructure:"OTEL_EXPORTER_OTLP_ENABLED" envconfig:"OTEL_EXPORTER_OTLP_ENABLED"`
}
