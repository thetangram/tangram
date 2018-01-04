package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	_, err := Load()
	if err != nil {
		t.Errorf("Error loading conf. Error: %v", err)
	}
}
