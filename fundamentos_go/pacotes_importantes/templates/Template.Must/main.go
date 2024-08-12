package main

import (
	"log"
	"os"
	"text/template"
)

type Pessoa struct {
	Nome  string
	Idade string
}

func main() {
	pessoa := Pessoa{Nome: "João", Idade: "30"}
	tmp := template.Must(template.New("PessoaTemplate").Parse("Nome: {{.Nome}} | Idade: {{.Idade}} anos"))
	err := tmp.Execute(os.Stdout, pessoa)
	if err != nil {
		log.Fatal(err)
	}

}
