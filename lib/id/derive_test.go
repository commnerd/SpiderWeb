package id

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
)

var parentId Id = Id(uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da"))

func TestDeriveBase(t *testing.T) {
	hexMap = []byte("0123456789abcdef")
	for i := 0; i <= 2; i++ {
		Derive(parentId, -1)
	}
	myId, _ := Derive(parentId, -1)

	assert.Equal(t, "3", string(myId.String()[0]))
}

func TestDeriveSub(t *testing.T) {
	hexMap = []byte("0123456789abcdef")
	Derive(parentId, 0)
	myId, _ := Derive(parentId, 0)

	assert.Equal(t, "31", string(myId.String()[0:2]))
}

func TestDeriveFromGoodMask(t *testing.T) {
	hexMap = []byte("0123456789abcdef")
	for i := 0; i <= 1; i++ {
		Derive(parentId, 7)
	}
	myId, _ := Derive(parentId, 7)

	assert.Equal(t, "322a1963-2", string(myId.String()[0:10]))
}

func TestDeriveFromBadMask(t *testing.T) {
	hexMap = []byte("0123456789abcdef")
	for i := 0; i <= 1; i++ {
		Derive(parentId, 7)
	}
	myId, _ := Derive(parentId, 8)

	assert.Equal(t, "322a1963-2", string(myId.String()[0:10]))
}