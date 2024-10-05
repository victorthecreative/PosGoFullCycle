package main

import "fmt"

func main() {
	ch := make(chan int)

	go calIdade(ch)
	reader(ch)
}

func calIdade(ch chan int) {
	anoBase := 2024
	anosNascimento := [5]int{1990, 2000, 2007, 1995, 2003}
	for i := range anosNascimento {
		idade := anoBase - anosNascimento[i]
		ch <- idade
	}
	close(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("A idade é: %d\n", x)
	}
}
