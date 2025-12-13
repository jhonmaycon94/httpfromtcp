package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error", "Error", err)
	}

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

			fmt.Printf("read: %s\n", line)
			line = ""
		}

		line += string(chunk)
	}

	if len(line) != 0 {
		fmt.Printf("read: %s\n", line)
	}
}
