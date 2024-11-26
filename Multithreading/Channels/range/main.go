package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	ch := make(chan int)

	go calIdade(ch)
	go reader(ch, &wg)

	wg.Wait()
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

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("A idade é: %d\n", x)
		wg.Done()
	}

}
