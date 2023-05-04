package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Carro andando..."
}

// Caso altere o nome da função, variável ou struct para minúsculo,
// soma, a ou carro [marcar], esses parâmetros não serão encontrados
// fora desse módulo!!!!!!
