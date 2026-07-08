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

// TestFileExistsPermissionDenied is a regression test for a bug where any
// os.Stat error other than "not exist" (e.g. permission denied on a parent
// directory) caused a nil-pointer panic because only os.IsNotExist(err) was
// checked before dereferencing stat.
func TestFileExistsPermissionDenied(t *testing.T) {
	if os.Getuid() == 0 {
		t.Skip("running as root, permission checks don't apply")
	}

	dir := t.TempDir()
	blockedFile := dir + "/inner/file.txt"

	require.NoError(t, os.Mkdir(dir+"/inner", 0755))
	require.NoError(t, os.Chmod(dir+"/inner", 0000))
	defer func() {
		_ = os.Chmod(dir+"/inner", 0755)
	}()

	assert.NotPanics(t, func() {
		assert.False(t, FileExists(blockedFile))
	})
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
