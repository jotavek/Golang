package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

//json.Marshal() converte um map ou um struct para JSON
//json.Unmarshal() converte um JSON em um map ou um struct

type usuario struct {
	Nome  string `json:"nome"`
	Idade uint   `json:"idade"`
}

func main() {
	user := usuario{"Joao", 22}

	userEmJson, erro := json.Marshal(user)
	if erro != nil {
		log.Fatal(erro)
	}
	// json.Marshal retorna um slice de bytes
	fmt.Println(userEmJson)
	// usando bytes.NewBuffer a gente tem o retorno esperado
	fmt.Println(bytes.NewBuffer(userEmJson))
}
