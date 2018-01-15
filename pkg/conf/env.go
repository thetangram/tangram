package conf

import (
	"log"
	"os"
	"time"
)

const (
	envAddress         = "TANGRAM_ADDRESS"
	envReadTimeout     = "TANGRAM_READ_TIMEOUT"
	envWriteTimeout    = "TANGRAM_WRITE_TIMEOUT"
	envShutdownTimeout = "TANGRAM_SHUTDOWN_TIMEOUT"
)

func (c *Config) loadEnv() {
	if env, exist := os.LookupEnv(envAddress); exist {
		c.addr = env
	}
	if env, exist := os.LookupEnv(envReadTimeout); exist {
		d, err := time.ParseDuration(env)
		if err == nil {
			c.readTimeout = d
		} else {
			log.Printf("Error parsing %v env variable. Expected duration but was %v\n", envReadTimeout, env)
		}
	}
	if env, exist := os.LookupEnv(envWriteTimeout); exist {
		d, err := time.ParseDuration(env)
		if err == nil {
			c.writeTimeout = d
		} else {
			log.Printf("Error parsing %v env variable. Expected duration but was %v\n", envWriteTimeout, env)
		}
	}
	if env, exist := os.LookupEnv(envShutdownTimeout); exist {
		d, err := time.ParseDuration(env)
		if err == nil {
			c.shutdownTimeout = d
		} else {
			log.Printf("Error parsing %v env variable. Expected duration but was %v\n", envShutdownTimeout, env)
		}
	}
}
