package main

import "fmt"

//Retorno nomeado voce coloca a variavel que vai ser retornada junto com o seu tipo
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(40, 35)
	fmt.Println(soma, subtracao)
}