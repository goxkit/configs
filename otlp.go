package configs

import "time"

type OTLPConfigs struct {
	//
	ExporterType string
	//
	Endpoint string
	//
	AccessKey string
	//
	ExporterTimeout time.Duration
	//
	ExporterInterval time.Duration
	//
	MetricsExporterRateBase float64
	//
	TracingExporterRateBase float64
	//
	MetricsEnabled bool
	//
	TracingEnabled bool
}
