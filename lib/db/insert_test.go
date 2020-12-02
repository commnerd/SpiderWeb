package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	record := &Record{
		Body: "blah blah blah",
	}
	respRecord, err := Insert(record)
	if err != nil {
		panic("Record exists")
	}
	assert.Equal(t, "blah blah blah", respRecord.Body)
}