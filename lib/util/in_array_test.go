package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
)

var intArray = [2, 4, 6, 8, 10]

func TestInArraySuccess(t *testing.T) {

	exists, index := InArray(4, intArray)

	assert.IsTrue(t, exists)
	assert.Equal(t, 1, index)
}

func TestInArrayFailure(t *testing.T) {

	exists, index := InArray(7, intArray)

	assert.IsTrue(t, !exists)
	assert.Equal(t, -1, index)
}