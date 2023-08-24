package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O client %s foi desativado", c.Nome)
}

// * interfaces
// qualquer Cliente que tive o método Desativar() impementará Pessoa
// somente para assinatura de métodos

type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()

}

func main() {
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 30,
		Ativo: true,
	}

	Desativacao(lucas)
}
