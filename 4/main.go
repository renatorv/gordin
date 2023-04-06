package main

import "fmt"

const a = "Hello!!"

type ID int

// DECLARAÇÃO DE VARIÁVEL
var (
	b bool
	c int     = 10
	d string  = "Renato"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de 'e' é %T, e seu valor e %v\n", e, e)
	fmt.Printf("O tipo de 'ID' é %T", f)
}
