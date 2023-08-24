package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	fmt.Printf("O cliente %v andou", c.nome)

}

type Conta struct {
	saldo int
}

func newConta() *Conta {
	return &Conta{saldo: 0}

}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {

	lucas := Cliente{
		nome: "Lucas",
	}

	conta1 := newConta()
	conta2 := newConta()

	fmt.Println(&conta1)
	fmt.Println(&conta2)

	conta1.simular(300)
	conta2.simular(550)

	fmt.Println(*conta1)
	fmt.Println(*conta2)

	lucas.andou()
}
