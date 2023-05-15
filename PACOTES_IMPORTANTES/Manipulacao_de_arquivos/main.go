package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// CRIAÇÃO
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	// GRAVAÇÃO
	// SE FOR SOMENTE UMA STRING: WriteString
	// tamanho, err := f.WriteString("Hello, World.")

	// SE FOR DADOS [BYTES]
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo."))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso!!\nTamanho: %d bytes.\n", tamanho)

	f.Close()

	// ABRIR O ARQUIVO PARA LEITURA
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// CASO O ARQUIVO SEJA MUITO GRANDE, PODE-SE FAZER A LEITURA DELE DE POUCO EM POUCO

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
