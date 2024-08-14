package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciado")
	defer log.Println("Request finalizado")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processado com sucesso!")
		w.Write([]byte("Request processado com sucesso!"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")                              //print command line stdout
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout) //print browser
	}

}
