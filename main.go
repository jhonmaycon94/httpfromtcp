package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(file io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer file.Close()
		defer close(ch)

		line := ""
		for {
			chunk := make([]byte, 8)
			i, err := file.Read(chunk)
			if err != nil {
				break
			}
			chunk = chunk[:i]
			if n := bytes.IndexByte(chunk, '\n'); n != -1 {
				line += string(chunk[:n])
				chunk = chunk[n+1:]
				ch <- line
				line = ""
			}

			line += string(chunk)
		}

		if len(line) != 0 {
			ch <- line
		}
	}()
	return ch
}

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error", "Error", err)
	}

	lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
