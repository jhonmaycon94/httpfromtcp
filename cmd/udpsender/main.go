package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("Erro: ", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Erro: ", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Erro: ", err)
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatal("Erro: ", err)
		}
	}
}
