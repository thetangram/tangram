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
	defaultConfFile        = "tangram.yml"
	defaultVerbose         = false
	defaultShowVersion     = false
	defaultHelp            = false
	defaultAddress         = ":2018"
	defaultReadTimeout     = 500 * time.Millisecond
	defaultWriteTimeout    = 2 * time.Second
	defaultShutdownTimeout = 5 * time.Second
)

// Config contains application configuration
type Config struct {
	// file is the configuration file
	file string
	// verbose starts tangram in verbose mode
	verbose bool
	// showVersion print version and exist
	showVersion bool
	// help prints in console the command line arguments and exists
	help            bool
	addr            string
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
	routes          []Route
}

// Route is a service routing configuration
type Route struct {
	path    string
	url     string
	timeout time.Duration
}

// Defaults creates a conf instance with default values.
func new() (c Config) {
	c = Config{
		file:            defaultConfFile,
		verbose:         defaultVerbose,
		showVersion:     defaultShowVersion,
		help:            defaultHelp,
		addr:            defaultAddress,
		readTimeout:     defaultReadTimeout,
		writeTimeout:    defaultWriteTimeout,
		shutdownTimeout: defaultShutdownTimeout,
		routes: []Route{
			Route{
				path:    "/hello-world",
				url:     "https://raw.githubusercontent.com/thetangram/hello-world/master/html/index.html",
				timeout: 5 * time.Second,
			},
		},
	}
	return
}

// Load application configuration.
// The configuration values can be set in environment variables,
// a configuration file and command line arguments. The way configuration
// is loaded is:
// 1.- First load environment variables
// 2.- Then load configuration file, and if some value overwrites the
//     values already stablished.
// 3.- Then reads the command line arguments, and overwrite values.
func Load() (c Config, err error) {
	c = new()
	c.loadEnv()
	c.loadCmd()
	return
}

// ShowVersion returns if version cmdline argument has been set
func (c Config) ShowVersion() bool {
	return c.showVersion
}

// Verbose returns if verbose cmdline argument has been set
func (c Config) Verbose() bool {
	return c.verbose
}

// Address gets the HTTP server address.
// The format is "[inet-address]:port"
func (c Config) Address() string {
	return c.addr
}

// ReadTimeout is the maximum duration for reading the entire request,
// including the body.
// See also https://golang.org/pkg/net/http/#Server ReadTimeout field
func (c Config) ReadTimeout() time.Duration {
	return c.readTimeout
}

// WriteTimeout is the maximum duration before timing out writes of
// the response.
// See also https://golang.org/pkg/net/http/#Server WriteTimeout field
func (c Config) WriteTimeout() time.Duration {
	return c.writeTimeout
}

// ShutdownTimeout gets the application shutdown timeout to wait to
// shutdown the HTTP server, for application graceful shutdown.
func (c Config) ShutdownTimeout() time.Duration {
	return c.shutdownTimeout
}

// Routes returns the array of configured routes.
func (c Config) Routes() []Route {
	return c.routes
}

// Path is the path Tangram routes to a service
func (r Route) Path() string {
	return r.path
}

// URL is the target address of a route
func (r Route) URL() string {
	return r.url
}

// Timeout is the max. time to perform request to service
func (r Route) Timeout() time.Duration {
	return r.timeout
}
