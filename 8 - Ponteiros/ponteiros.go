package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referencia de memoria
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Print(variavel3, *ponteiro)
	//Quando eu coloco um * antes do ponteiro acontece uma desreferenciação

	
}
