package main

import "fmt"

func main() {
	fmt.Println("main")
	fmt.Println(sum(1, 334, 535, 3534, 2, 645, 25, 3646, 57, 567, 2525, 546))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total

}
