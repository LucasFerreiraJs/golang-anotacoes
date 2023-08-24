package main

// só temos for mesmo
func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	stringSlice := []string{"teste", "teste2", "teste3", "teste4"}

	for k, v := range stringSlice {
		print(k, "-", v, "\n")
	}

	i := 0
	for i < 10 {
		print(i)
		print("\n")
		i++
	}
}
