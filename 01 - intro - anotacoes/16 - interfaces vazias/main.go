package main

import "fmt"

// * geral que tiver o método Y irá implementar a interface Teste
type Teste interface {
	y()
}

// * geral implementa essa interface (por estar vazia)
// cuidado
type TesteVazia interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "teste"

	showType(x)
	showType(y)
}

func showType(value interface{}) {

	fmt.Printf("O tipo da variável é %T e o valor é %v\n", value, value)

}
