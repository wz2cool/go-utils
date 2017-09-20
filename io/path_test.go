package io

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestFileOrDirExists(t *testing.T) {
	tempDir := os.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")
	os.Remove(testFile)
	isExist, err := FileOrDirExists(testFile)
	if err != nil {
		t.Error(err)
	}

	if isExist {
		t.Error("this file should not exists")
	}

	log.Println(testFile)
	// create file for testing.
	fs, err := os.Create(testFile)
	if err != nil {
		t.Error(err)
	}
	defer fs.Close()
	isExist, err = FileOrDirExists(testFile)
	if err != nil {
		t.Error(err)
	}

	if !isExist {
		t.Error("this file should exists")
	}
}
