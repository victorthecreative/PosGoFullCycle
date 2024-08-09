package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		resp, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o Body: %v\n", err)
		}

		var infocep ViaCep
		err = json.Unmarshal(resp, &infocep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter Json: %v\n", err)
		}
		file, err := os.Create("Cidade.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", infocep.Cep, infocep.Localidade, infocep.Uf))
	}

}
