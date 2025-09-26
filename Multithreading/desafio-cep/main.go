package main

import (
	"cep_desafio/cmd"
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	cmd.Execute()
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf(" Buscando o CEP  \r%c", r)
			time.Sleep(delay)
		}
	}

}
