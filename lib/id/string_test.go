package id

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
)

func TestString(t *testing.T) {
	myString := New().String()

	myUUID := uuid.MustParse(myString).String();

	assert.Equal(t, myString, myUUID)
}