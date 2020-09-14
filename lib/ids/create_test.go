package ids

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
)

func TestCreateIDType(t *testing.T) {
	myId := CreateID()

	assert.IsType(t, uuid.New(), myId);
}
