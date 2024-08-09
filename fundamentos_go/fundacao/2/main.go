package main

import "fmt"

func main() {
	var numeros [3]int
	numeros[0] = 1
	numeros[1] = 2
	numeros[2] = 3

	// ultimo := numeros[len(numeros)-1]

	// fmt.Println(ultimo)

	for _, v := range numeros {
		fmt.Println(v)
	}

}
