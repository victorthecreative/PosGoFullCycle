package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, WaitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Taks %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		WaitGroup.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	go task("A", &waitGroup)
	go task("B", &waitGroup)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Taks %s is running\n", i, "Anonima")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
