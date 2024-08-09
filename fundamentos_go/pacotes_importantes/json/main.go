package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero  int    `json:"N"`
	Saldo   int    `json:"S"`
	Titular string `josn:"T"`
}

func MarshalStruct(conta Conta) string {
	tr, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	byte_to_string := string(tr)

	return byte_to_string
}

func EncoderStruct(conta Conta) {
	err := json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}
}

func main() {

	// conta := Conta{Numero: 253654, Saldo: 5432, Titular: "Victor"}
	// marshal_Json := MarshalStruct(conta)
	// fmt.Println(marshal_Json)

	// EncoderStruct(conta)

	var conta2 Conta
	Json := []byte(`{"N":8796,"S":356,"T":"Julia"}`)

	err := json.Unmarshal(Json, &conta2)
	if err != nil {
		panic(err)
	}

	println(conta2.Titular)

}
