package main

// DECLARAÇÃO DE VARIÁVEL
var (
	b bool
	c int     = 10
	d string  = "Renato"
	e float64 = 1.2
)

const a = "Hello!!"

func main() {
	// var x string = "x" // variavel local deve ser utilizada
	x := "x" // string
	x = "xx"
	b = true
	println(b)
	println(c)
	println(d)
	println(e)
	println(x)
}
