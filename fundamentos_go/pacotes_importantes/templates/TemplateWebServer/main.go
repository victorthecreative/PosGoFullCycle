package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Curso struct {
	Nome    string
	Duracao int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(os.Stdout, Cursos{
			{"Pós Go", 300},
			{"Pós Python", 200},
			{"Pós Java", 400},
		})
		if err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":8182", nil)

}
