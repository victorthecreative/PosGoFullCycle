package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d esta recebendo %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {

	data := make(chan int)
	QtdWorker := 10

	for i := 0; i < QtdWorker; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}

}
