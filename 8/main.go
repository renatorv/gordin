package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(sum3(51, 2))
	valor, err := sum3(5, 10)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

// ****************************************************
// caso o valor de a e b fossem do msm tipo, pode-se *
// realizar a entrada de parâmetros da seuinte forma *
// func sum(a, b int) int {                          *
// ****************************************************
func sum(a int, b int) int {
	return a + b
}

// ****************************************************
//
//	FUNÇÕES EM GO PODEM RETORNAR DOIS VALORES!!    *
//
// ****************************************************
func sum2(a int, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

// ****************************************************
// GO não tem try catch, então usa-se, por padrão o  *
// parâmetro error para fazer a validação se a função*
// executou corretamente ou não, conforme abaixo.    *
// ****************************************************
// nil: significa vazio (nulo)                       *
// ****************************************************
func sum3(a int, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
