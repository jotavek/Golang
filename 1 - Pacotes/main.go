package main

import (
	"fmt"
	"modulo/auxiliar"
)

// Função escreve uma mensagem e chama a função Escrever do package auxiliar
func main() {
	fmt.Println("Escrevendo o arquivo main")
	auxiliar.Escrever()
}
