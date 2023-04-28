package main

func soma(a, b int) int {
	a = 50
	return a + b
}

func somaViaPonteiro(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {

	minhaVar1 := 10
	minhaVar2 := 20
	soma(minhaVar1, minhaVar2)
	println(minhaVar1)
	println(minhaVar2)

	varPonteiro1 := 100
	varPonteiro2 := 200
	somaViaPonteiro(&varPonteiro1, &varPonteiro2)
	println(varPonteiro1)
	println(varPonteiro2)
}
