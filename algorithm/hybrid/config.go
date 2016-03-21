package hybrid

import (
	"runtime"
)

// Config represents a configuration of the algorithm.
type Config struct {
	MinLevel uint // Minimal level of interpolation
	MaxLevel uint // Maximum level of interpolation

	TotalError float64 // Tolerance on the total error
	LocalError float64 // Tolerance on the local error

	Workers uint // Number of concurrent workers
}

// NewConfig returns a new configuration with default values.
func NewConfig() *Config {
	return &Config{
		MinLevel: 1,
		MaxLevel: 10,

		TotalError: 1e-6,
		LocalError: 1e-6,

		Workers: uint(runtime.GOMAXPROCS(0)),
	}
}
