package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {

	lucas := Client{
		Nome:  "Lucas",
		Idade: 18,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, idade %d, Ativo: %t", lucas.Nome, lucas.Idade, lucas.Ativo)
}
