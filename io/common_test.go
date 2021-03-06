package io

import (
	"os"
	"path/filepath"
	"testing"
)

var tempDir = os.TempDir()
var testFile = filepath.Join(tempDir, "testFile.txt")
var testParentDir = filepath.Join(tempDir, "testParentDir")
var testDir = filepath.Join(testParentDir, "testDir")

func deferRemoveTestFile(t *testing.T) {
	err := removeTestFile()
	if err != nil {
		t.Error(err)
	}
}

func deferRemoveTestDir(t *testing.T) {
	err := removeTestDir()
	if err != nil {
		t.Error(err)
	}
}

func createTestFile() error {
	fs, err := os.Create(testFile)
	defer fs.Close()
	return err
}

func removeTestFile() error {
	// create file again, the simplest way to avoid file not found issue.
	err := createTestFile()
	if err != nil {
		return err
	}

	return os.Remove(testFile)
}

func createTestDir() error {
	return os.MkdirAll(testDir, 0711)
}

func removeTestDir() error {
	return os.RemoveAll(testParentDir)
}

func clearnTestDir() error {
	err := removeTestDir()
	if err != nil {
		return err
	}

	err = createTestDir()
	if err != nil {
		return err
	}
	return nil
}
