package main

import "fmt"

func main() {
	var minhaVar interface{} = "Wesley Willians"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o valor de ok é %v", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor de red2 é %v", res2)
}
