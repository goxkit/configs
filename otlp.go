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
	Endpoint string `mapstructure:"OTLP_ENDPOINT" envconfig:"OTLP_ENDPOINT"`
	//ENV: OTLP_ACCESS_KEY
	//
	// AccessKey is the access key for authentication with the OTLP endpoint
	AccessKey string `mapstructure:"OTLP_ACCESS_KEY" envconfig:"OTLP_ACCESS_KEY"`
	//ENV: OTLP_EXPORTER_TIMEOUT
	//
	// ExporterTimeout specifies the timeout duration for the OTLP exporter
	ExporterTimeout time.Duration `mapstructure:"OTLP_EXPORTER_TIMEOUT" envconfig:"OTLP_EXPORTER_TIMEOUT"`
	//ENV: OTLP_EXPORTER_INTERVAL
	//
	// ExporterInterval specifies the interval duration for the OTLP exporter
	ExporterInterval time.Duration `mapstructure:"OTLP_EXPORTER_INTERVAL" envconfig:"OTLP_EXPORTER_INTERVAL"`
	//ENV: OTLP_METRICS_EXPORTER_RATE_BASE
	//
	// MetricsExporterRateBase specifies the rate base for metrics exporter
	MetricsExporterRateBase float64 `mapstructure:"OTLP_METRICS_EXPORTER_RATE_BASE" envconfig:"OTLP_METRICS_EXPORTER_RATE_BASE"`
	//ENV: OTLP_TRACING_EXPORTER_RATE_BASE
	//
	// TracingExporterRateBase specifies the rate base for tracing exporter
	TracingExporterRateBase float64 `mapstructure:"OTLP_TRACING_EXPORTER_RATE_BASE" envconfig:"OTLP_TRACING_EXPORTER_RATE_BASE"`
	// ENV: OTLP_METRICS_ENABLED
	//
	// MetricsEnabled indicates whether metrics collection is enabled
	MetricsEnabled bool `mapstructure:"OTLP_METRICS_ENABLED" envconfig:"OTLP_METRICS_ENABLED"`
	// ENV: OTLP_TRACING_ENABLED
	//
	// TracingEnabled indicates whether tracing collection is enabled
	TracingEnabled bool `mapstructure:"OTLP_TRACING_ENABLED" envconfig:"OTLP_TRACING_ENABLED"`
}
