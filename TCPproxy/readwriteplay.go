package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type barReader struct{}

func (b *barReader) Read(p []byte) (int, error) {
	fmt.Print("in>")
	return os.Stdin.Read(p)
}

type barWriter struct{}

func (b *barWriter) Write(p []byte) (int, error) {
	fmt.Print("out>")
	return os.Stdout.Write(p)
}

func main() {
	var (
		reader barReader
		writer barWriter
	)
	// buff := make([]byte, 5025)
	// byt, err := reader.Read(buff)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("we read %b from standard in\n", byt)

	// wbyt, err := writer.Write(buff)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("we write %b to standard output\n", wbyt)
	w, err := io.Copy(&writer, &reader)
	if err != nil {
		log.Fatal("not copied")
	}
	fmt.Printf("we write %b to standard output\n", w)
}
