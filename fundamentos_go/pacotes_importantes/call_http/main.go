package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://google.com.br")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	res_string := string(res)

	fmt.Println(res_string)

}
