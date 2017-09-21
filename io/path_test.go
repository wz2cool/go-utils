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

func TestFileExists(t *testing.T) {
	err := createTestFile()
	defer deferRemoveTestFile(t)

	if err != nil {
		t.Error(err)
		return
	}
	exists := FileOrDirExists(testFile)

	if !exists {
		t.Error("file should exists")
		return
	}
}

func TestFileNotExists(t *testing.T) {
	err := removeTestFile()
	if err != nil {
		t.Error(err)
		return
	}

	exists := FileOrDirExists(testFile)
	if err != nil {
		t.Error(err)
		return
	}

	if exists {
		t.Error("file should not exists")
		return
	}
}

func TestDirExists(t *testing.T) {
	err := createTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	defer deferRemoveTestDir(t)

	exists := FileOrDirExists(testDir)
	if !exists {
		t.Errorf("Dir: %s; should exists", testDir)
		return
	}
}

func TestDirNotExists(t *testing.T) {
	err := removeTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	exists := FileOrDirExists(testDir)
	if exists {
		t.Errorf("Dir: %s; should not exists", testDir)
		return
	}
}

func TestCreateDirIfNotExists(t *testing.T) {
	// clean dir first
	err := removeTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	err = CreateDirIfNotExists(testDir)
	if err != nil {
		t.Error(err)
		return
	}

	exists := FileOrDirExists(testDir)
	if !exists {
		t.Errorf("Dir: %s, should exists", testDir)
		return
	}
}

func TestNotCreateDirIfExists(t *testing.T) {
	err := createTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	defer deferRemoveTestDir(t)
	err = CreateDirIfNotExists(testDir)
	if err != nil {
		t.Error(err)
		return
	}
}

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
