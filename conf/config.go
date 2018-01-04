// Deal with application configuration.
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

type Config struct {
	addr string
}

// get the HTTP server address
func (c Config) Address() string {
	return c.addr
}

// Loads applcation configuration from default config file
func Load() (c Config, err error) {
	return loadConfig(defaultConfigFile)
}

func loadConfig(file string) (c Config, err error) {
	c = Config{}
	c.addr = defaultAddress
	return
}
