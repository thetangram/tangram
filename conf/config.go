// Package config models the application configuration.
// The configuration values will be loaded from command arguments,
// config file, environment variables and default values.
// The prevalence order is (more to less prevalence):
//   - command arguments
//   - config file
//   - environment
//   - default values
package config

const (
	defaultConfigFile = "tangram.toml"
	// defaults
	defaultAddress = ":2018"
)

// Config contains application configuration
type Config struct {
	addr string
}

// Address returns the HTTP server address.
// This format have the format "[inet-address]:port"
func (c Config) Address() string {
	return c.addr
}

// Load applcation configuration from default config file
func Load() (c Config, err error) {
	return loadConfig(defaultConfigFile)
}

func loadConfig(file string) (c Config, err error) {
	c = Config{}
	c.addr = defaultAddress
	return
}
