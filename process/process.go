// Package process provides helper functions for process operations.
package process

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetExecutableName returns the filename of the execution binary.
// The complete path is omitted. Only the filename is returned.
func GetExecutableName() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("could not get executable name: %w", err)
	}

	return filepath.Base(path), nil
}
