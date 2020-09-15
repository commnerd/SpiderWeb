package ids

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"testing"
)

var parentId uuid.UUID = uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")

func TestCreateTopLevelId(t *testing.T) {
	myId := Create(parentId, 0)

	assert.Equal(t, "3", string(myId.String()[0]))
}
