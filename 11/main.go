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

// Isso significa que qq struct que tiver o método Desativar
// está implementado a interface Pessoa automaticamente.
// Ou seja qq cliente que tenha o método Desativar vai implementar a
// interface Pessoa.
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // composicao
	// ou
	Address Endereco // criando uma propriedade do tipo Endereco
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	// renato := Cliente{
	// 	Nome:  "Renato",
	// 	Idade: 45,
	// 	Ativo: true,
	// }

	// renato.Ativo = false
	// renato.Cidade = "Rio Verde"
	// renato.Desativar()

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)

	// fmt.Printf("\nNome: %s, Idade: %d, Ativo: %t\n", renato.Nome, renato.Idade, renato.Ativo)
}
