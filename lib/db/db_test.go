package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbInstanceCreation(t *testing.T) {
	db := New();
	assert.IsType(t, &DB{}, db)
	assert.Equal(t, 0, len(db.Collections))
}