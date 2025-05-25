// Copyright (c) 2025, The GoKit Authors
// MIT License
// All rights reserved.

package configs

// HTTPConfigs provides configuration settings for HTTP servers and clients.
// It contains parameters related to network binding, addressing, and diagnostic features.
type HTTPConfigs struct {
	//ENV: HTTP_ENABLE_PROFILING
	//
	// EnableProfiling controls whether Go's pprof profiling endpoints should be exposed
	// for runtime debugging and performance analysis
	EnableProfiling bool `mapstructure:"HTTP_ENABLE_PROFILING" envconfig:"HTTP_ENABLE_PROFILING"`
}

const (
	HTTPEnableProfilingEnvKey = "HTTP_ENABLE_PROFILING" // Whether to enable pprof profiling endpoints
)
