package main

func main() {
	a := 10 // Memória -> Endereço -> Valor

	// Endereço de memória
	println(&a)

	// Ponteiro: é o endereçamento da memória
	// [ variável -> ponteiro que tem um endereço na memória -> valor ]
	// '*' indica que estamos tratando do endereçamento de memória
	var ponteiro *int = &a
	println(ponteiro)

	*ponteiro = 20

	println(a)

	b := &a

	println(b)
	println(*b)
	println(&b)

}
