package main

import (
	"fmt"
)

func main() {

	total := func() int {
		return sum(51, 2, 13, 5, 96, 84, 754, 236, 4, 99, 56, 6) * 2
	}()

	fmt.Printf("Retorno da closure: %v\n", total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Printf("Retorno de sum(): %v\n", total)
	return total
}
