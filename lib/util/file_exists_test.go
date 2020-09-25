package util

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"fmt"
	"os"
)

func TestMain(m *testing.M) {
	file := []byte("This file is under test.")
    err := ioutil.WriteFile("./file_under_test.txt", file, 0644)
	check(err)

	m.Run()

	err = os.Remove("./file_under_test.txt")
	check(err)
}

func TestFileExists(t *testing.T) {
	assert.True(t, FileExists("./file_under_test.txt"))
}

func TestFileNotExists(t *testing.T) {
	assert.True(t, !FileExists("./foo_bar.txt"))
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}