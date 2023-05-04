package main

import (
	"fmt"

	"github.com/renatorv/gordin/matematica"

	"github.com/google/uuid"
)

type MyNumber int

func main() {
	res := matematica.Soma(10, 20)

	fmt.Println("Resultado:", res)

	fmt.Println(matematica.A)

	carro := matematica.Carro{Marca: "Volkwagem"}
	fmt.Println(carro.Marca)
	fmt.Println(carro.Andar())

	fmt.Println(uuid.New())

}

// go mod init <nome_modulo>
// go mod init github.com/renatorv/gordin
// nome_modulo: sugestão endereço do github
// github.com/renatorv/curso-go

// INSTALANDO MÓDULOS
// go get github.com/google/uuid
// ou
// github.com/google/uuid no import e depois: go mod tidy no terminal
