package main

import "fmt"

//Estrutura de uma função anônima
func main() {
	func() {
		fmt.Println("Olá Mundo!")
	}()

	//Função Anônima com parâmetros.

	func(texto string) {
		fmt.Println(texto)
	}("Olá Mundo!")

	//Função Anônima com return
	//fmt.sprintf, %s é para string %d para numeros

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 4)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
