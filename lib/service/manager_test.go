package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
	"../ids"
)

type mockNode struct {}

func (node *mockNode) GetId() uuid.UUID {
	return uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")
}

func (node *mockNode) GetMask() ids.Mask {
	return ids.Mask(2)
}

func TestNewManager(t *testing.T) {
	manager := NewManager(&mockNode{})

	assert.Equal(t, len(manager.Services), 0)
	assert.IsType(t, NewManager(&mockNode{}), manager)
}