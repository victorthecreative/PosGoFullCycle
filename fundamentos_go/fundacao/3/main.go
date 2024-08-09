package main

import "fmt"

func main() {
	numeros_pares := []int{2, 4, 6, 8}
	fmt.Printf("O tamnho da lista é %d, e sua capacidade é %d\n", len(numeros_pares), cap(numeros_pares))

	numeros_pares = append(numeros_pares, 10)

	fmt.Printf("Foi adicionado um novo item, agora o tamanho é %d e sua dobrou a capacidade para %d", len(numeros_pares), cap(numeros_pares))
}
