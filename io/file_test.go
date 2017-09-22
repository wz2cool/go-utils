package io

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	err := createTestFile()
	if err != nil {
		t.Error(err)
		return
	}

	err = clearnTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	defer deferRemoveTestDir(t)

	distPath := filepath.Join(testDir, "test.txt")
	err = CopyFile(testFile, distPath)
	if err != nil {
		t.Error(err)
		return
	}

	exists := FileOrDirExists(distPath)
	if !exists {
		t.Error("file should exists")
		return
	}
}

func TestGetMD5(t *testing.T) {
	err := removeTestFile()
	if err != nil {
		t.Error(err)
		return
	}

	expectedMD5 := "3e25960a79dbc69b674cd4ec67a72c62"
	err = createFileForTestingMD5(testFile, "Hello world")
	if err != nil {
		t.Error(err)
		return
	}
	defer removeTestFile()

	actucalMD5, err := GetFileMD5(testFile)
	if err != nil {
		t.Error(err)
		return
	}

	if expectedMD5 != actucalMD5 {
		t.Error("md5 should be equal!")
		return
	}
}

func TestGetMD5FileNotFound(t *testing.T) {
	err := removeTestFile()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = GetFileMD5(testFile)
	if err == nil {
		t.Error("should have error")
		return
	}
}

func createFileForTestingMD5(filePath, content string) error {
	fileDir := filepath.Dir(filePath)
	err := CreateDirIfNotExists(fileDir)
	if err != nil {
		return err
	}

	fs, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fs.Close()
	_, err = fs.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
