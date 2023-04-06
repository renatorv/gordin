package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(51, 2, 13, 5, 96, 84, 754, 236, 4, 99, 56, 6))
}

// *****************************************************
//
//	FUNÇÕES VARIÁTICAS                    *
//
// Recebem um número variável de parâmetros e processa *
// esses valores através de um loop.                   *
// *****************************************************
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
