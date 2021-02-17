package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var intArray = []int{2, 4, 6, 8, 10}

func TestInArraySuccess(t *testing.T) {

	exists, index := InArray(4, intArray)

	assert.True(t, exists)
	assert.Equal(t, 1, index)
}

func TestInArrayFailure(t *testing.T) {

	exists, index := InArray(7, intArray)

	assert.True(t, !exists)
	assert.Equal(t, -1, index)
}
