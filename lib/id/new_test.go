package id

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
)

func TestNew(t *testing.T) {
	myId := New()

	assert.IsType(t, Id(uuid.New()), myId)
}