package global

import (
	"runtime"
)

// Config represents a configuration of the algorithm.
type Config struct {
	MaxLevel uint // Maximum level of interpolation

	AbsoluteError float64 // Tolerance on the absolute error
	RelativeError float64 // Tolerance on the relative error

	Workers uint // Number of concurrent workers
}

// NewConfig returns a new configuration with default values.
func NewConfig() *Config {
	return &Config{
		MaxLevel: 10,

		AbsoluteError: 1e-6,
		RelativeError: 1e-3,

		Workers: uint(runtime.GOMAXPROCS(0)),
	}
}
