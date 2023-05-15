package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {
		// println(cep)
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		defer req.Body.Close()

		// LER O CONTEÚDO DO BODY
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		// res é um Json, assim vamos transforma-ló em uma Struct
		// JSON {
		// 	"cep": "01001-000",
		// 	"logradouro": "Praça da Sé",
		// 	"complemento": "lado ímpar",
		// 	"bairro": "Sé",
		// 	"localidade": "São Paulo",
		// 	"uf": "SP",
		// 	"ibge": "3550308",
		// 	"gia": "1004",
		// 	"ddd": "11",
		// 	"siafi": "7107"
		//   }
		// Esse site gera a Struct com base em um Json: https://mholt.github.io/json-to-go/

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
		}

		// Gravando no arquivo:
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar aquivo: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))

		fmt.Println("Arquivo criado com sucesso!!")
		fmt.Println(data.Localidade)
	}
}
