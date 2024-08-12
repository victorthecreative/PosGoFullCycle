package main

import (
	"html/template"
	"log"
	"os"
)

type Curso struct {
	Nome      string
	TempoHora int
}

func main() {
	curso := Curso{"Pós Go", 300}
	tmp := template.New("CursoTemplate")

	tmp, _ = tmp.Parse("Nome: {{.Nome}} | Tempo: {{.TempoHora}} horas")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		log.Fatal(err)
	}

}
