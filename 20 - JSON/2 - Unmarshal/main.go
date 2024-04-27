package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// json.Marshal retorna um slice de bytes e o json.Unmarshal precisa receber um slice de bytes
// json.Unmarshal([]byte(os dados a serem passados), endereço de memória onde vai ser passado os dados é necessario usar &)
type user struct {
	Nome  string `json:"nome"`
	Idade uint   `json:"idade"`
}

func main() {
	userEmJson := `{"nome":"Joao","idade":22}`

	var usuario user
	// Tem que usar o parametro com & para representar um endereço de memória e também porque eu quero que a variável usuario fique alterado
	if erro := json.Unmarshal([]byte(userEmJson), &usuario); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(usuario)
}
