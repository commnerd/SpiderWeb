package node

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
)

var parent Node = New(nil)

func TestNew(t *testing.T) {
	parent.Id = uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")
	child := New(&parent)

	assert.Equal(t, "3", string(child.Id.String()[0]))
	assert.Equal(t, 0, child.Mask)
	assert.Equal(t, child.Parent, &parent)
}
