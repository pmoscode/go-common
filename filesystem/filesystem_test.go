package filesystem

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileExists(t *testing.T) {
	file := createTempFile(t)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Printf("could not remove temp file: %s, error: %v", name, err)
		}
	}(file)

	assert.True(t, FileExists(file), "expected file to exist: %s", file)
}

func TestFileNotExists(t *testing.T) {
	file := "not-existent-file.txt"

	assert.False(t, FileExists(file), "file should not exist: %s", file)
}

func createTempFile(t *testing.T) string {
	t.Helper()
	f, err := os.CreateTemp("./", "temp-*.txt")
	require.NoError(t, err, "could not create temp file")
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Printf("could not close temp file: %s, error: %v", f.Name(), err)
		}
	}(f)

	return f.Name()
}
