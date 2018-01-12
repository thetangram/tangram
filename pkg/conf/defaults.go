package conf

import (
	"time"
)

const (
	defaultAddress         = ":2018"
	defaultReadTimeout     = 500 * time.Millisecond
	defaultWriteTimeout    = 2 * time.Second
	defaultShutdownTimeout = 5 * time.Second
)

// Defaults creates a conf instance with default values
func defaults() (c Config, err error) {
	c = Config{
		addr:            defaultAddress,
		readTimeout:     defaultReadTimeout,
		writeTimeout:    defaultWriteTimeout,
		shutdownTimeout: defaultShutdownTimeout,
	}
	return
}
