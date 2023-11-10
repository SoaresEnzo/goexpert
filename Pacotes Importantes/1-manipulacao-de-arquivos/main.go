package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Hello World!")
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %v bytes\n", tamanho)
	f.Close()
	// Leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		tamanho, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:tamanho]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
