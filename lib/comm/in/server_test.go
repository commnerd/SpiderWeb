package in

import (
	"github.com/stretchr/testify/assert"
	// "github.com/google/uuid"
	"testing"
	// "os"
)

// func TestMain(m *testing.M) {
// 	parent = New(nil)
// 	parent.Id = uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")
//     code := m.Run()
//     os.Exit(code)
// }

func TestNewServer(t *testing.T) {
	s := NewServer()

	assert.IsType(t, &server{}, s)
}