package main

import (
	"fmt"

	"github.com/renatorv/gordin/matematica"
)

type MyNumber int

func main() {
	res := matematica.Soma(10, 20)

	fmt.Println("Resultado:", res)

}

// go mod init <nome_modulo>
// go mod init github.com/renatorv/gordin
// nome_modulo: sugestão endereço do github
// github.com/renatorv/curso-go
