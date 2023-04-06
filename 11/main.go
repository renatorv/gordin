package main

import "fmt"

//**************************************************
// GO não é O.O.                                   *
// O que mais se aproxima de assemlha as classes,  *
// métodos no GO são as Structs                    *
//**************************************************
// GO não tem herança, porém pode-se trabalhar com *
// composição de várias structs                    *
//**************************************************

type Endereco struct {
	Logradouro string
	Numero     string
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // composicao
	// ou
	Address Endereco // criando uma propriedade do tipo Endereco
}

func main() {
	renato := Cliente{
		Nome:  "Renato",
		Idade: 45,
		Ativo: true,
	}

	renato.Ativo = false
	renato.Cidade = "Rio Verde"

	fmt.Printf("\nNome: %s, Idade: %d, Ativo: %t\n", renato.Nome, renato.Idade, renato.Ativo)
}

next => Métodos em Structs
