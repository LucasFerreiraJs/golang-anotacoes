package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("hello world")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes", tamanho)

	f.Close()
	// * Leitura
	// não precisa dar close
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	println("\n")
	fmt.Println(arquivo)
	println("\n")
	fmt.Println(string(arquivo))

	// nesse caso precisamos de um buffer
	outroArquivo, err := os.Open("arquivo2.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(outroArquivo)
	buffer := make([]byte, 10) // slice de 10 bytes

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	outroArquivo.Close()
}
