package main

import (
	"encoding/json"
	"os"
)

// Tags: são anotações que podem ser colocadas nas structs no formato: chave/valor
type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	// Marshal: Transforma[Serializa] em Json
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	// A transformação em json SEMPRE retorna em bytes
	println(string(res))

	// Se já se desejar fazer algo com o resultado, pode-se criar um
	// encoder, que irá atribuir o resultado para: saída padrão(nesse ex),
	// web server e etc
	err = json.NewEncoder(os.Stdout).Encode(conta)
	// Aqui se der um erro ele será atribuido aqui na variável err.
	if err != nil {
		println(err)
	}

	// TRANSFORMANDO UM JSON EM UMA CONTA:
	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		print(err)
	}
	println(contaX.Saldo)

	jsonPuro2 := []byte(`{"n":3,"s":300}`)
	var contaX2 Conta
	err = json.Unmarshal(jsonPuro2, &contaX2)
	if err != nil {
		print(err)
	}
	println(contaX2.Saldo)

}
