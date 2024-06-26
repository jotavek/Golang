package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

// Função escreve uma mensagem e chama a função Escrever do package auxiliar
func main() {
	fmt.Println("Escrevendo o arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("jota@gmail.com")
	fmt.Println(erro)
}
