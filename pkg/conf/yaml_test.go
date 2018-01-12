package conf

import (
	"testing"
	"time"
)

var (
	yamlContent = []byte(`
address: :2018
read-timeout: 5s
write-timeout: 5s
shutdown-timeout: 20s
routes:
- path: /one
  url: http://component-1
  timeout: 2s
- path: /two
  url: http://component-2
  timeout: 4s
`)
)

// TestConfigFile check that a yaml structure can be unmarshaled
func TestUnmarshal(t *testing.T) {
	address := ":2018"
	readTimeout := 5 * time.Second
	writeTimeout := 5 * time.Second
	shutdownTimeout := 20 * time.Second
	routes := 2
	routeOnePath := "/one"

	c := ConfigYAML{}
	err := c.unmarshal(yamlContent)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if c.Addr != address {
		t.Fatalf("error parsing address: Expected: [%v], Returned: [%v]\n", address, c.Addr)
	}
	if c.ReadTimeout != readTimeout {
		t.Fatalf("error parsing read timeout: Expected: [%v], Returned: [%v]\n", readTimeout, c.ReadTimeout)
	}
	if c.WriteTimeout != writeTimeout {
		t.Fatalf("error parsing read timeout: Expected: [%v], Returned: [%v]\n", writeTimeout, c.WriteTimeout)
	}
	if c.ShutdownTimeout != shutdownTimeout {
		t.Fatalf("error parsing read timeout: Expected: [%v], Returned: [%v]\n", shutdownTimeout, c.ShutdownTimeout)
	}
	if len(c.Routes) != routes {
		t.Fatalf("error parsing routes: Expected: [%v] routes, Returned: [%v] routes\n", routes, len(c.Routes))
	}
	if c.Routes[0].Path != routeOnePath {
		t.Fatalf("error parsing routes: Expected: [%v] routes, Returned: [%v] routes\n", routeOnePath, c.Routes[0].Path)
	}
}

// TestConfigFile check that a yaml structure can be unmarshaled
func TestToConfig(t *testing.T) {
	configYAML := ConfigYAML{
		Addr:         ":2017",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	c, err := configYAML.toConfig()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if c.addr != configYAML.Addr {
		t.Fatalf("error parsing address: Expected: [%v], Returned: [%v]\n", configYAML.Addr, c.addr)
	}
	if c.readTimeout != configYAML.ReadTimeout {
		t.Fatalf("error parsing read timeout: Expected: [%v], Returned: [%v]\n", configYAML.ReadTimeout, c.readTimeout)
	}
	if c.writeTimeout != configYAML.WriteTimeout {
		t.Fatalf("error parsing read timeout: Expected: [%v], Returned: [%v]\n", configYAML.WriteTimeout, c.writeTimeout)
	}
}

// TestToRoutes check that a ConfigYAML routes array is adapted to Config routes array
func TestToRoutes(t *testing.T) {
	configYAML := ConfigYAML{
		Routes: []RouteYAML{
			RouteYAML{
				Path:    "/path",
				URL:     "http://domain.com/path/to/component",
				Timeout: 3 * time.Second,
			},
			RouteYAML{
				Path:    "/path",
				URL:     "http://domain.com/path/to/component",
				Timeout: 3 * time.Second,
			},
		},
	}

	r := configYAML.toRoutes()
	if len(r) != 2 {
		t.Fatalf("error parsing routes: Expected: [%v], Returned: [%v]\n", 2, len(r))
	}
}

// TestToRoute check that a yaml route structure is adapted to Route
func TestToRoute(t *testing.T) {
	routeYAML := RouteYAML{
		Path:    "/path",
		URL:     "http://domain.com/path/to/component",
		Timeout: 3 * time.Second,
	}

	r := routeYAML.toRoute()
	if r.path != routeYAML.Path {
		t.Fatalf("error parsing path: Expected: [%v], Returned: [%v]\n", routeYAML.Path, r.path)
	}
	if r.url != routeYAML.URL {
		t.Fatalf("error parsing url: Expected: [%v], Returned: [%v]\n", routeYAML.URL, r.url)
	}
	if r.timeout != routeYAML.Timeout {
		t.Fatalf("error parsing timeout: Expected: [%v], Returned: [%v]\n", routeYAML.Timeout, r.timeout)
	}
}
