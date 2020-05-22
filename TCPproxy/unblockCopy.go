package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type fooReader struct{}

func (b *fooReader) Read(p []byte) (int, error) {
	fmt.Print("in>")
	return os.Stdin.Read(p)
}

type fooWriter struct{}

func (b *fooWriter) Write(p []byte) (int, error) {
	fmt.Print("out>")
	return os.Stdout.Write(p)
}

func main() {
	var (
		reader fooReader
		writer fooWriter
	)

	//Run in gorountine to prevent copy from blocking
	go func() {
		if _, err := io.Copy(&writer, &reader); err != nil {
			log.Fatalln("not copied")
		}
	}()
	file, err := os.OpenFile("copyfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Failed to create or open file")
	}
	if _, err := io.Copy(file, &reader); err != nil {
		log.Fatalln("failed to write to file")
	}

}
