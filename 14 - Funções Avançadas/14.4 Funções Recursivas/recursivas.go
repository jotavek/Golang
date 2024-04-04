package main

import "fmt"

//Função recursiva vai chamar ela mesma
//Ela nao é utilizado para casos que tem muitas execuções.

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	posicao := uint(6)
	fmt.Println(fibonacci(posicao))

}
