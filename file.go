package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	conteudo, err := os.ReadFile("document.txt")
	if err != nil {
		fmt.Println("Ops, esse arquivo não existe!", err)
		os.Exit(1)
	}

	texto := strings.TrimSpace(string(conteudo))

	palavras := strings.Split(texto, ",")

	for i, palavra := range palavras {
		palavra = strings.TrimSpace(palavra)
		fmt.Printf("A palavra %d é: %s\n", i+1, palavra)
	}
}
