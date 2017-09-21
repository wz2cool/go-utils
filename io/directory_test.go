package io

import (
	"os"
	"path/filepath"
	"testing"
)

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

func TestDirIsEmpty(t *testing.T) {
	err := removeTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	err = createTestDir()
	if err != nil {
		t.Error(err)
		return
	}
	defer deferRemoveTestDir(t)

	isEmpty, err := IsDirEmpty(testDir)
	if err != nil {
		t.Error(err)
		return
	}

	if !isEmpty {
		t.Error("dir should be empty")
		return
	}
}

func TestDirIsNotEmpty(t *testing.T) {
	err := createTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	defer deferRemoveTestDir(t)

	testFileInDir := filepath.Join(testDir, "testFileInDir.txt")
	fs, err := os.Create(testFileInDir)
	if err != nil {
		t.Error(err)
		return
	}
	fs.Close()

	isEmpty, err := IsDirEmpty(testDir)
	if err != nil {
		t.Error(err)
		return
	}

	if isEmpty {
		t.Error("dir should not be empty")
		return
	}
}

func TestIsDirEmptyPathNotExists(t *testing.T) {
	err := removeTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = IsDirEmpty(testDir)
	if err != nil {
		t.Logf("expected err: %s", err)
		return
	}
	t.Error("should not found path")
}

func TestIsDirEmptyFileCanotOpen(t *testing.T) {
	err := createTestFile()
	if err != nil {
		t.Error(err)
		return
	}

	defer deferRemoveTestFile(t)
	_, err = IsDirEmpty(testFile)
	if err != nil {
		t.Logf("expected err: %s", err)
		return
	}
	t.Error("should not found path")
}
