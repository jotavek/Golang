package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 4 * 3
	restodadivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restodadivisao)

	// ATRIBUICAO
	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM DOS OPERADORES RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("-----------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //and
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++     //incrementação
	numero += 15 // numero = numero + 15
	numero--     // decrementação
	numero -= 20 // numero = numero - 20

	numero *= 3 // numero = numero * 3
	fmt.Println(numero)
	// FIM DOS OPERADORES UNÁRIOS

	// OPERADOR TERNÁRIO
	// Nao é possivel em go
	// texto := numero > 5 ? "Maior que 5" : "Menor que 5"
	// Mas é possivel com if e else
	var texto string
	if numero > 40 {
		texto = "Maior que 40"
	} else {
		texto = "Menor que 40"
	}
	fmt.Println(texto)
}
