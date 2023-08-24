package main

// extends/implements ?
type MyNumber int

type CustomNumber interface {
	~int | ~float64
}

// ~
// considera type do tipo

func Soma[T CustomNumber](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

// * [T Type] GENERICS
// func Compara[T CustomNumber](a T, b T) bool {

// algumas coisas não são comparable
// * pra isso tem o package constrainsts

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{
		"Wesley": 100,
		"joao":   2000,
		"Maria":  4000,
	}

	m2 := map[string]float64{
		"Wesley": 100.30,
		"joao":   2000.60,
		"Maria":  3000.40,
	}
	m3 := map[string]MyNumber{
		"Wesley": 100,
		"joao":   2000,
		"Maria":  3000,
	}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println("\n")

	println(Compara(10, 3))
}
