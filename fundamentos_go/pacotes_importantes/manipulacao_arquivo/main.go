package main

import (
	"bufio"
	"fmt"
	"os"
)

func create_and_write() {
	crt, err := os.Create("teste.txt")
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo: %s", err)
	}

	crt.WriteString("Ola, teste!")

	crt.Close()
}

func create_and_writ_binary() {
	crt, err := os.Create("teste_bytes.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := crt.Write([]byte("Ola, teste!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	crt.Close()
}

func read_to_archive() {
	arquivo, err := os.ReadFile("teste.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))
}

func read_large_archive() {
	read, err := os.Open("Teste.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(read)
	buffer := make([]byte, 3)
	for {
		N, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:N]))
	}

}

func main() {
	// create_and_write()
	// create_and_writ_binary()
	// read_to_archive()
	read_large_archive()
}
