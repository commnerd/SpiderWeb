package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestInstanceAddCollection(t *testing.T) {
	db := New();
	err := db.AddCollection("foo")
	assert.True(t, err == nil)
	assert.IsType(t, &Collection{}, db.Collections["foo"])
}

func TestAddCollection(t *testing.T) {
	AddCollection("foo")
	assert.IsType(t, &Collection{}, instance.Collections["foo"])
	instance = New()
}

func TestFailureOnRepeatCollection(t *testing.T) {
	db := New();

	err := db.AddCollection("foo")
	assert.True(t, err == nil)

	err = db.AddCollection("foo")
	assert.True(t, err != nil)
	assert.Equal(t, err.Error(), fmt.Sprintf(CollectionPresentError, "foo"))
}