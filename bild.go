//This is the entrypoint for the frontend of bild.

package main

import (
	"bild/pipeline"
	"bild/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	fmt.Println("Bild is Screenshot as a Service")
	args := os.Args[1:] //only get the actual arguments

	//TODO handle input better, now we expect
	filePath := args[0]
	fmt.Println("File with urls is", filePath)
	start(filePath)
}

func start(filePath string) {
	var wg sync.WaitGroup
	file := util.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//each url is sent to a goroutine.
		wg.Add(1)
		go pipeline.Pipeline(scanner.Text(), &wg)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Something went wrong while reading the file.", err)
	}

	wg.Wait()
	fmt.Println("Bild is done...!")
}
