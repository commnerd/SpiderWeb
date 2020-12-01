package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstanceInsert(t *testing.T) {
	record := &Record{ Body: "blah blah blah" }
	respInt := instance.Insert(record)
	resp := respInt.(*Record)
	assert.Equal(t, "blah blah blah", resp.Body)
}