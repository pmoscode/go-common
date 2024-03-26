// Package filesystem provides some useful functions handling files and directories
package filesystem

import (
	"os"
)

// FileExists checks, if a file exists.
// Returns true, if so or false, if the file does not exist.
// It also returns false, if the file is a directory.
func FileExists(filename string) bool {
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !stat.IsDir()
}
