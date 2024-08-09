package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	CepParam := r.URL.Query().Get("cep")
	if CepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, error := BuscaCep(CepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCep, error) {
	resp, erro := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if erro != nil {
		return nil, erro
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c ViaCep

	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}

	return &c, nil
}
