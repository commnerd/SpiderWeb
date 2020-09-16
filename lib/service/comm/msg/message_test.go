package msg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoinRequestLabel(t *testing.T) {

	assert.Equal(t, JoinRequest.String(), "Join Request")
}