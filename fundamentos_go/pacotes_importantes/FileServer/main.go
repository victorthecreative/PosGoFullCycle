package main

import (
	"log"
	"net/http"
)

func main() {

	FileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", FileServer)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
