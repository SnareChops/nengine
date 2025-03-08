package config_test

import (
	"strings"
	"testing"

	"github.com/SnareChops/nengine/config"
	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	A int    `yaml:"A"`
	B string `yaml:"B"`
}

func TestConfigLoad(t *testing.T) {
	def := TestConfig{A: 1, B: "test"}
	reader := strings.NewReader("B: other")
	c, err := config.NewConfig(reader, def)
	assert.Nil(t, err)
	c.Load()
	f := c.Get()
	assert.Equal(t, 1, f.A)
	assert.Equal(t, "other", f.B)
}
