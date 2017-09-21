package io

import (
	"testing"
)

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
