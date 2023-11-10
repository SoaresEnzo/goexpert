package main

import (
	"fmt"

	"curso-go/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Andar())
	fmt.Println(carro)
	fmt.Printf("Resultado: %v\n", s)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}
