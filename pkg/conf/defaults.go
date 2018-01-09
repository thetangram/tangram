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

func defaults() (c Config, err error) {
	c = Config{
		addr: defaultAddress,
		http: http{
			readTimeout:  defaultHTTPReadTimeout,
			writeTimeout: defaultHTTPWriteTimeout,
		},
		system: system{
			shutdownTimeout: defaultSystemShutdownTimeout,
		},
	}
	return
}
