package util

import (
	"os"
)

func FileExists(file string) bool {
	info, err := os.Stat(file)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}