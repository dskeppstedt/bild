package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

//OpenFile creats a file descriptor for the given filepath.
func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file")
		return nil
	}
	return file
}

func WriteFile(filepath string, data []byte) {
	err := ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		fmt.Println("Failed to write file", err)
	}
}
