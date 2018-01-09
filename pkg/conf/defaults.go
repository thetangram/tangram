package conf

import (
	"time"
)

const (
	defaultAddress         = ":2018"
	defaultShutdownTimeout = 5 * time.Second
)

func defaults() (c Config, err error) {
	c = Config{
		addr:            defaultAddress,
		shutdownTimeout: defaultShutdownTimeout}
	return
}
