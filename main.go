package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
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
	listener, err := net.Listen("tcp", "127.0.0.1:42069")
	if err != nil {
		log.Fatal("Error", "Error", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connection accepted")

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Println(line)
		}
	}
}
