package node

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"../config"
)

func TestConfigSetting(t *testing.T) {
	assert.True(t, config.Get("NODE") != nil)
}