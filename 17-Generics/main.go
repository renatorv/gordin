package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	si := map[string]int{"Renato": 1000, "Neide": 2000, "Neuza": 3000}
	sf := map[string]float64{"Renato": 1000.0, "Neide": 2000.0, "Neuza": 4000.0}

	sn := map[string]MyNumber{"Renato": 1000, "Neide": 2000, "Neuza": 3000}

	println(Soma(si))
	println(Soma(sf))
	println(Soma(sn))

	println(Compara(10, 10))

}
