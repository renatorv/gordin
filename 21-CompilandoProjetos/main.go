// go build main.go
// GOOS=linux go build main.go
// go env
// https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures

// Caso o app tenha um módulo, o nome do executável será o nome do módulo.
// Para mudar isso: go build -o xpto

package main

func main() {
	a := 1
	b := 2
	c := 3

	// || ou
	// && e
	if a > b || c > a {
		println("a > b || c > a")
	}

	if a > b {
		println("a")
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}

}
