// Package util defines and implements helper methods
package util

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// LeftPad pads a string with a padStr for a certain number of times
func LeftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}

// RightPadToLen pads a string to a certain length
func RightPadToLen(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	retStr := s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// FileExists determines if the named file exists
func FileExists(filePath string) bool {

	f, err := os.Open(filePath)
	f.Close()
	if err != nil {
		return false
	}
	return true
}

// CurrentDirectory returns the current directory
func CurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// CopyFile copies the file from the source to the destination file
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			os.Chmod(dest, sourceinfo.Mode())
		}
	}

	return
}
