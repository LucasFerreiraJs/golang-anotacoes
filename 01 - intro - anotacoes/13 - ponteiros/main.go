package main

func main() {

	// * memória -> endereco ->  valor
	a := 10

	// variável aponta pra um ponteiro
	println(a)
	println(&a) // -> end na memória

	var bPonteiro *int = &a
	*bPonteiro = 50

	//mais fácil
	b := &a
	*b = 4
	println(a)
	println(b)
	println(*b)

}
