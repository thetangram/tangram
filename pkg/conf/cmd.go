package conf

import (
	"flag"
)

func (c *Config) loadCmd() {
	flag.StringVar(&c.addr, "address", c.addr, "The address to bind the HTTP listener. Default: "+defaultAddress)
	flag.DurationVar(&c.readTimeout, "read-timeout", c.readTimeout, "Read timeout for incoming HTTP request. Default: "+defaultReadTimeout.String())
	flag.DurationVar(&c.writeTimeout, "write-timeout", c.writeTimeout, "HTTP read timeout. Default: "+defaultWriteTimeout.String())
	flag.DurationVar(&c.shutdownTimeout, "shutdown-timeout", c.shutdownTimeout, "HTTP read timeout. Default: "+defaultShutdownTimeout.String())
	flag.StringVar(&c.file, "file", c.file, "The configuration file. Default: "+defaultConfFile)
	flag.BoolVar(&c.verbose, "verbose", c.verbose, "Verbose mode")
	flag.BoolVar(&c.showVersion, "version", c.showVersion, "Shows version")
	flag.BoolVar(&c.help, "help", c.help, "Shows help and exits")
	flag.Parse()
	return
}
