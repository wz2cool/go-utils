package io

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
)

// CopyFile copys source file to destination
func CopyFile(source, dest string) error {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err != nil {
		return err
	}

	// try to set same mode.
	si, err := os.Stat(source)
	if err == nil {
		os.Chmod(dest, si.Mode())
	}
	return nil
}

const filechunk = 8192 // we settle for 8KB
// GetFileMD5 returns md5 hash of file
func GetFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// calculate the file size
	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	filesize := info.Size()
	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))
	md5Hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		if _, err = io.WriteString(md5Hash, string(buf)); err != nil {
			return "", err
		} // append into the hash
	}
	return fmt.Sprintf("%x", md5Hash.Sum([]byte(""))), nil
}
