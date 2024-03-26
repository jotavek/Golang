package main

import (
	"errors"
	"fmt"
)

func main() {
	// NUMEROS INTEIROS
	var numero int = 1000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 3000
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// FIM NUMEROS INTEIROS

	//n NUMEROS REAIS
	var numeroReal1 float32 = 123.44
	fmt.Println(numeroReal1)
	var numeroReal2 float64 = 1212.4324344
	fmt.Println(numeroReal2)

	//Inferencia de tipo pega a arquitetura do seu PC como base, no meu caso foi 64.
	numeroReal3 := 312.4932490321
	fmt.Println(numeroReal3)

	// FIM NUMEROS REAIS

	// STRINGS
	var str string = "blabla"
	fmt.Println(str)

	str2 := "blabla2"
	fmt.Println(str2)

	char := "A"
	fmt.Println(char)

	// FIM STRINGS

	//

	var texto int16
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
