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

// Config contains application configuration
type Config struct {
	addr   string
	http   http
	system system
}

type http struct {
	readTimeout  time.Duration
	writeTimeout time.Duration
}

type system struct {
	shutdownTimeout time.Duration
}

// Load application configuration
func Load() (c Config, err error) {
	c, _ = defaults()
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
	return c.system.shutdownTimeout
}

// HTTPReadTimeout is the maximum duration for reading the entire request, 
// including the body.
// See also https://golang.org/pkg/net/http/#Server ReadTimeout field
func (c Config) HTTPReadTimeout() time.Duration {
	return c.http.readTimeout
}

// HTTPWriteTimeout is the maximum duration before timing out writes of
// the response. 
// See also https://golang.org/pkg/net/http/#Server WriteTimeout field
func (c Config) HTTPWriteTimeout() time.Duration {
	return c.http.writeTimeout
}
