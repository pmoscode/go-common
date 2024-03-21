package filesystem

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	file := createTempFile()
	defer os.Remove(file)

	if !FileExists(file) {
		listDirectory()
		t.Fatal("Expected file not found: ", file)
	}
}
func TestFileNotExists(t *testing.T) {
	file := "not-existent-file.txt"

	if FileExists(file) {
		listDirectory()
		t.Fatal("File should not exist: ", file)
	}
}

func createTempFile() string {
	dir := "./" // replace with your desired directory
	pattern := "temp-*.txt"

	f, err := os.CreateTemp(dir, pattern)
	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
	}
	defer f.Close()

	return f.Name()
}

func listDirectory() {
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
