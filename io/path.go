package io

import (
	"os"
)

// FileOrDirExists returns a boolean describing if the path of file or directory exists.
func FileOrDirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	isExist := os.IsExist(err)
	return isExist
}
