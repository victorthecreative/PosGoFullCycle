package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(sum(3, 4))
	value, err := sum_two_retunrs(25, 30)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

func sum_two_retunrs(d, e int) (int, error) {
	if d+e >= 50 {
		return 0, errors.New("Erro")
	}
	return d + e, nil
}
