package port

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	p := Port(1234)
	assert.Equal(t, "1234", p.String())
}