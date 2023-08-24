package main

func soma(a, b int) int {
	return a + b
}

// recebendo valores reais
func soma2(a, b *int) int {
	*a = 40
	return *a + *b
}
func main() {

	value1 := 10
	value2 := 8

	println(soma(value1, value2))
	println(soma2(&value1, &value2))
	println(value1)

	//
}
