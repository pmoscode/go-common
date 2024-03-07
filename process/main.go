package process

import (
	"os"
	"path/filepath"
)

func GetExecutableName() string {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Base(path)
}
