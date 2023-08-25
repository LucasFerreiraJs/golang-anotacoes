package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"` // talvez seria comparado com anotations... talvez
	Saldo  int `json:"s"`
}

func main() {

	conta := Conta{Numero: 1, Saldo: 100}

	// * transformar em json
	// sempre retorna em bytes
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(res) // bytes
	println("\n")
	println(string(res))

	// encoder
	// já transforma em json entregando para "alguém"
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{
		"Numero" : 2,
		"Saldo" : 200
		}`)

	var novaConta Conta
	err = json.Unmarshal(jsonPuro, &novaConta)
	if err != nil {
		panic(err)
	}

	println(novaConta.Saldo)

	novoJson := []byte(`{
		"n": 3,
		"s": 550
		}`)

	var novaConta2 Conta
	err = json.Unmarshal(novoJson, &novaConta2)
	if err != nil {
		println(err)
	}

	println(novaConta2.Saldo)
}
