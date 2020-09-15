package node

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
	"../ids"
)

var parent Node = New(nil)

func TestNew(t *testing.T) {
	parent.Id = uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")
	child := New(&parent)

	assert.Equal(t, "3", string(child.Id.String()[0]))
	assert.Equal(t, ids.Mask(0), child.Mask)
	assert.Equal(t, child.Parent, &parent)
}

func TestGetNextBadMask(t *testing.T) {
	parent.Id = uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")
	parent.Mask = ids.Mask(7)
	child := New(&parent)

	assert.Equal(t, "322a1963-2", string(child.Id.String()[0:child.Mask+1]))
	assert.Equal(t, ids.Mask(9), child.Mask)
	assert.Equal(t, child.Parent, &parent)
}

func TestGetNextMask(t *testing.T) {
	parent.Id = uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")
	parent.Mask = ids.Mask(26)
	child := New(&parent)

	assert.Equal(t, "322a1963-2b7f-43d4-b9cf-2fce", string(child.Id.String()[0:child.Mask+1]))
	assert.Equal(t, ids.Mask(27), child.Mask)
	assert.Equal(t, child.Parent, &parent)
}
