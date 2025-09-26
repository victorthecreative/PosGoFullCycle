package services

import (
	"cep_desafio/app/domain"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	viaCepURL    = "https://viacep.com.br/ws/%d/json/"
	brasilApiURL = "https://brasilapi.com.br/api/cep/v1/%d"
)

func GetCEPInfo(cep int64, servico string, ch chan interface{}) {

	var viacep *domain.ViaCep
	var brasilapi *domain.BrasilAPI

	if servico == "viacep" {
		r, err := http.Get(fmt.Sprintf(viaCepURL, cep))
		defer r.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		viacep = domain.NewViacep()
		err = json.Unmarshal(body, &viacep)
		if err != nil {
			log.Fatal(err)
		}
		ch <- viacep
	} else if servico == "brasilapi" {
		r, err := http.Get(fmt.Sprintf(brasilApiURL, cep))
		defer r.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		brasilapi = domain.NewBrasilAPI()
		err = json.Unmarshal(body, &brasilapi)
		if err != nil {
			log.Fatal(err)
		}
		ch <- brasilapi
	}
}
