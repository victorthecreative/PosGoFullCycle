package main

import "fmt"

func main() {

	canal := make(chan string)

	go func() {
		canal <- "Ola mundo"
	}()

	msg := <-canal

	fmt.Println(msg)
}
