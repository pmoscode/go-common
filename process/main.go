// Package process provides helper functions for process operations.
package process

import (
	"os"
	"path/filepath"
)

// GetExecutableName returns the filename of the execution binary.
// The complete path is omitted. Only the filename is returned.
func GetExecutableName() string {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Base(path)
}
