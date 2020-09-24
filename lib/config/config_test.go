package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/spf13/viper"
	"testing"
)

func TestGetConfig(t *testing.T) {
	assert.IsType(t, viper.New(), GetConfig())
}