package main

import (
	"io"
	"os"
)

func main() {
	filename := os.Args[1]

	filePointer, _ := os.Open(filename)

	io.Copy(os.Stdout, filePointer)
}
