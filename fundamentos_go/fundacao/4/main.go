package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "Victor": 2000, "Maria": 3000}

	fmt.Println(salarios["Victor"])
}
