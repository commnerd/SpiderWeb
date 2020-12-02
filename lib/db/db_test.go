package db

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"../config"
)

func TestNew(t *testing.T) {
	db := New()
	assert.IsType(t, &DB{}, db)
	assert.Equal(t, 0, len(db.Collections))
}

func TestGet(t *testing.T) {
	local := Get()
	assert.Equal(t, instance, local)
}

func TestSetInConfig(t *testing.T) {
	db, ok := config.Get("DB").(*DB)
	assert.True(t, ok)
	assert.Equal(t, instance, db)
}