package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 8, 5, 9, 9, 5, 4, 2, 6))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
