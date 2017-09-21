package io

import (
	"io"
	"os"
)

// CreateDirIfNotExists creates a direcotry named path,
// if direcotry not exists, then create it or do nothing.
func CreateDirIfNotExists(path string) error {
	if _, err := os.Stat(path); err == nil {
		// path exists
		return nil
	}
	// path not exists
	err := os.MkdirAll(path, 0711)
	return err
}

// IsDirEmpty returns a boolean describing if directory is empty.
func IsDirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
