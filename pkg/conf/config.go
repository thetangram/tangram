// Package conf models the application configuration.
// The configuration values will be loaded from command line arguments,
// config file, environment variables and default values.
// The order is (more to less prevalence):
// - command arguments
// - config file
// - environment
// - default values
package conf

import (
	"time"
)

const (
	defaultAddress         = ":2018"
	defaultShutdownTimeout = 5 * time.Second
)

// Config contains application configuration
type Config struct {
	addr            string
	shutdownTimeout time.Duration
}

// Load application configuration
func Load() (c Config, err error) {
	c = Config{
		addr:            defaultAddress,
		shutdownTimeout: defaultShutdownTimeout}
	return
}

// Address gets the HTTP server address.
// The format is "[inet-address]:port"
func (c Config) Address() string {
	return c.addr
}

// ShutdownTimeout gets the application shutdown timeout to wait to
// shutdown the HTTP server, for application graceful shutdown.
func (c Config) ShutdownTimeout() time.Duration {
	return c.shutdownTimeout
}
