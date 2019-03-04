package util

import (
	"log"
	"os"
)

//OpenFile creats a file descriptor for the given filepath.
func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
