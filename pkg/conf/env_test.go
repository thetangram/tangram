package conf

import (
	"os"
	"testing"
	"time"
)

// TestNoEnvVariablesAndNothingIsChangedInConfig verifies that when
// no env variables, nothing changes.
func TestNoEnvVariablesAndNothingIsChangedInConfig(t *testing.T) {
	// sut
	c := Config{}
	expected := Config{}
	c.loadEnv()
	// assertions
	if c.addr != expected.addr {
		t.Fail()
	}
	if c.readTimeout != expected.readTimeout {
		t.Fail()
	}
	if c.writeTimeout != expected.writeTimeout {
		t.Fail()
	}
	if c.shutdownTimeout != expected.shutdownTimeout {
		t.Fail()
	}
}

// TestNoEnvVariablesAndNothingIsChangedInConfig verifies that when
// no env variables, nothing changes.
func TestLoadConfigFromEnv(t *testing.T) {
	expectedAddress := ":2019"
	expectedReadTimeout := 10 * time.Second
	expectedWriteTimeout := 20 * time.Second
	expectedShutdownTimeout := 30 * time.Second

	os.Setenv(envAddress, expectedAddress)
	os.Setenv(envReadTimeout, expectedReadTimeout.String())
	os.Setenv(envWriteTimeout, expectedWriteTimeout.String())
	os.Setenv(envShutdownTimeout, expectedShutdownTimeout.String())

	// sut
	c := Config{}
	c.loadEnv()
	// assertions
	if c.addr != expectedAddress {
		t.Fail()
	}
	if c.readTimeout != expectedReadTimeout {
		t.Fail()
	}
	if c.writeTimeout != expectedWriteTimeout {
		t.Fail()
	}
	if c.shutdownTimeout != expectedShutdownTimeout {
		t.Fail()
	}
}

// TestNoEnvVariablesAndNothingIsChangedInConfig verifies that when
// no env variables, nothing changes.
func TestEnvWithWrongDurationFormat(t *testing.T) {
	os.Setenv(envAddress, defaultAddress)
	os.Setenv(envReadTimeout, "wrong")
	os.Setenv(envWriteTimeout, "duration")
	os.Setenv(envShutdownTimeout, "format")

	// sut
	println()
	c := Config{}
	c.loadEnv()
	// assertions
	if c.addr != defaultAddress {
		t.Errorf("[%v] [%v]\n", c.addr, defaultAddress)
	}
	if c.readTimeout != 0 {
		t.Errorf("[%v] [%v]\n", c.readTimeout, defaultReadTimeout)
	}
	if c.writeTimeout != 0 {
		t.Errorf("[%v] [%v]\n", c.writeTimeout, defaultWriteTimeout)
	}
	if c.shutdownTimeout != 0 {
		t.Errorf("[%v] [%v]\n", c.shutdownTimeout, defaultShutdownTimeout)
	}
}
