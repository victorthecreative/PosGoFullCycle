package main

import (
	"net/http"
	"time"
)

func main() {
	C := http.Client{Timeout: time.Microsecond}
	resp, err := C.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
