package io

import "os"

// FileOrDirExists returns a boolean describing if the path of file or directory exists.
func FileOrDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// CreateDirIfNotExists creates a direcotry named path,
// if direcotry not exists, then create it or do nothing.
func CreateDirIfNotExists(path string) error {
	if _, err := os.Stat(path); err == nil {
		// path exists
	} else {
		// path not exists
		err := os.MkdirAll(path, 0711)
		if err != nil {
			return err
		}
	}
	return nil
}
