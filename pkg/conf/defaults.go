package conf

import (
	"time"
)

const (
	defaultAddress               = ":2018"
	defaultHTTPReadTimeout       = 500 * time.Millisecond
	defaultHTTPWriteTimeout      = 2 * time.Second
	defaultSystemShutdownTimeout = 5 * time.Second
)

// Defaults creates a conf instance with default values
func defaults() (c Config, err error) {
	c = Config{
		http: http{
			addr:         defaultAddress,
			readTimeout:  defaultHTTPReadTimeout,
			writeTimeout: defaultHTTPWriteTimeout,
		},
		system: system{
			shutdownTimeout: defaultSystemShutdownTimeout,
		},
	}
	return
}
