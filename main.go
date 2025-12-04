package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err == nil {
		message := make([]byte, 8)
		i, err := file.Read(message)
		if err == nil {
			fmt.Println(string(message[:i]))
		} else {
			fmt.Printf("Erro ao ler arquivo: %s", err)
		}

	} else {
		fmt.Printf("erro ao abrir o arquivo: %s", err)
	}
}
