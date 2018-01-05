// Package config models the application configuration.
// The configuration values will be loaded from command arguments,
// config file, environment variables and default values.
// The prevalence order is (more to less prevalence):
//   - command arguments
//   - config file
//   - environment
//   - default values
package config

import (
    "os"
)

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
	c.addr = confValue("address", "server.address", "ADDRESS", defaultAddress)
	return
}

// confValue returns a configuration value from best-fit source.
// Parameters:
// - arg: the name of configuration value in the command line argument 
// - conf: the name of configuration value in configuration file 
// - env: the name of configuration value in the environment 
// - def: the default value 
func confValue(arg string, conf string, env string, def string) string {
    if value, exist := os.LookupEnv(env); exist {
        return value
    }
    return def
}
