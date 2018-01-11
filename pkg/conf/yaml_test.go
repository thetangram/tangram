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

	address         = ":2018"
	readTimeout     = 5 * time.Second
	writeTimeout    = 5 * time.Second
	shutdownTimeout = 20 * time.Second
)

// TestConfigFile check that a yaml structure can be unmarshaled
func TestUnmarshal(t *testing.T) {
	c := ConfigYAML{}
	err := c.unmarshal(yamlContent)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if address != c.Addr {
		t.Fatalf("error parsing address: Expected: [%v], Returned: [%v]\n", address, c.Addr)
	}
	if readTimeout != c.ReadTimeout {
		t.Fatalf("error parsing read timeout: Expected: [%v], Returned: [%v]\n", readTimeout, c.ReadTimeout)
	}
	if writeTimeout != c.WriteTimeout {
		t.Fatalf("error parsing read timeout: Expected: [%v], Returned: [%v]\n", writeTimeout, c.WriteTimeout)
	}
	if shutdownTimeout != c.ShutdownTimeout {
		t.Fatalf("error parsing read timeout: Expected: [%v], Returned: [%v]\n", shutdownTimeout, c.ShutdownTimeout)
	}
	if len(c.Routes) != 2 {
		t.Fatalf("error parsing routes: Expected: [%v] routes, Returned: [%v] routes\n", 2, len(c.Routes))
	}
	if c.Routes[0].Path != "/one" {
		t.Fatalf("error parsing routes: Expected: [%v] routes, Returned: [%v] routes\n", "/one", c.Routes[0].Path)
	}
}

// TestConfigFile check that a yaml structure can be unmarshaled
func TestToConfig(t *testing.T) {
	configYAML := ConfigYAML{}
	c, err := configYAML.toConfig()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if configYAML.Addr != c.addr {
		t.Fatalf("error parsing address: Expected: [%v], Returned: [%v]\n", configYAML.Addr, c.addr)
	}
	if configYAML.ReadTimeout != c.readTimeout {
		t.Fatalf("error parsing read timeout: Expected: [%v], Returned: [%v]\n", configYAML.ReadTimeout, c.readTimeout)
	}
}
