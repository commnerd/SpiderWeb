package message_bus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	assert.IsType(t, &bus{}, instance)
}