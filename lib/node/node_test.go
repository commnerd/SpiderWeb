package node

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
	"../ids"
	"os"
)

var parent *Node

func TestMain(m *testing.M) {
	parent = New(nil)
	parent.Id = uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")
    code := m.Run()
    os.Exit(code)
}

func TestNew(t *testing.T) {
	child := New(parent)

	assert.Equal(t, "3", string(child.Id.String()[0]))
	assert.Equal(t, ids.Mask(0), child.Mask)
	assert.Equal(t, child.Parent, parent)
}

func TestGetNextBadMask(t *testing.T) {
	parent.Mask = ids.Mask(7)
	child := New(parent)

	assert.Equal(t, "322a1963-2", string(child.Id.String()[0:child.Mask+1]))
	assert.Equal(t, ids.Mask(9), child.Mask)
	assert.Equal(t, child.Parent, parent)
}

func TestGetNextMask(t *testing.T) {
	parent.Mask = ids.Mask(26)
	child := New(parent)

	assert.Equal(t, "322a1963-2b7f-43d4-b9cf-2fce", string(child.Id.String()[0:child.Mask+1]))
	assert.Equal(t, ids.Mask(27), child.Mask)
	assert.Equal(t, child.Parent, parent)
}

func TestAddChild(t *testing.T) {
	parent.Mask = ids.Mask(1)
	child := New(parent);
	parent.AddChild(child);

	assert.Equal(t, ids.Mask(2), child.Mask)
	assert.Equal(t, "32", string(child.Id.String()[0:child.Mask]))
	assert.Equal(t, child, parent.Children[child.Id.String()])
}