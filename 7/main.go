package main

import "fmt"

func main() {

	salarios := map[string]int{"Renato ": 100, "Ligero ": 200, "Gabriel": 300}
	// fmt.Println(salarios["Renato"])
	// delete(salarios, "Renato")
	salarios["Ren    "] = 500
	// fmt.Println(salarios["Ren"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Renato"] = 100
	// sal["Ligero"] = 500

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	// for _, salario := range salarios {
	// 	fmt.Printf("O salário é %d\n", salario)
	// }

}
