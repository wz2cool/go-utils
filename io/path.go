package io

import "os"

/// FileOrDirExists checking file
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

/// CreateDirIfNotExists
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
