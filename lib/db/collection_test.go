package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestCollectionNotExists(t *testing.T) {
	db := New();
	exists := db.CollectionExists("foo")
	assert.True(t, !exists)
}

func TestCollectionExists(t *testing.T) {
	db := New();
	db.Collections["foo"] = &Collection{}
	exists := db.CollectionExists("foo")
	assert.True(t, exists)
}

func TestAddCollection(t *testing.T) {
	db := New();
	err := db.AddCollection("foo")
	assert.True(t, err == nil)
	assert.IsType(t, &Collection{}, db.Collections["foo"])

	// cleanup
	instance = New()
}

func TestInstanceAddCollection(t *testing.T) {
	AddCollection("foo")
	assert.IsType(t, &Collection{}, instance.Collections["foo"])

	// cleanup
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

func TestFailureOnInstanceRepeatCollection(t *testing.T) {
	err := instance.AddCollection("foo")
	assert.True(t, err == nil)

	err = instance.AddCollection("foo")
	assert.True(t, err != nil)
	assert.Equal(t, err.Error(), fmt.Sprintf(CollectionPresentError, "foo"))

	// cleanup
	instance = New()
}