package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Renato Alves"
	fmt.Printf("\nO cliente %v andou\n", c.nome)
}

func main() {

	renato := Cliente{nome: "Renato"}

	renato.andou()

	fmt.Printf("O valor da struct com nome %v", renato.nome)

}
