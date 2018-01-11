package conf

import (
	"time"

	"github.com/go-yaml/yaml"
)

// ConfigYAML is a system configuration in the YAML file
type ConfigYAML struct {
	Addr            string        `yaml:"address"`
	ReadTimeout     time.Duration `yaml:"read-timeout"`
	WriteTimeout    time.Duration `yaml:"write-timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown-timeout"`
	Routes          []RouteYAML   `yaml:"routes"`
}

// RouteYAML is the service routing configuration in the YAML file
type RouteYAML struct {
	Path    string        `yaml:"path"`
	URL     string        `yaml:"url"`
	Timeout time.Duration `yaml:"timeout"`
}

func loadFromYAML(b []byte) (c Config, err error) {
	loaded := ConfigYAML{}
	err = loaded.unmarshal(b)
	if err != nil {
		return
	}
	c, err = loaded.toConfig()
	if err != nil {
		return
	}
	return
}

func (c *ConfigYAML) unmarshal(b []byte) (err error) {
	err = yaml.Unmarshal(b, &c)
	return
}

func (conf *ConfigYAML) toConfig() (c Config, err error) {
	c.addr = conf.Addr
	c.readTimeout = conf.ReadTimeout
	c.writeTimeout = conf.WriteTimeout
	c.shutdownTimeout = conf.ShutdownTimeout
	c.routes = conf.toRoutes()
	return
}

func (conf *ConfigYAML) toRoutes() (r []Route) {
	r = make([]Route, len(conf.Routes))
	for i := 0; i < len(r); i++ {
		r[0] = conf.Routes[i].toRoute()
	}
	return
}

func (route *RouteYAML) toRoute() (r Route) {
	r.path = route.Path
	r.url = route.URL
	r.timeout = route.Timeout
	return
}
