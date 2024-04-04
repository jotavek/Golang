package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")

}
func funcao2() {
	fmt.Println("Executando a função 2")

}

// O defer é executado imediatamente antes do retorno da função.
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Mostrando Resultado logo abaixo.")
	media := (n1 + n2) / 2

	if media >= 5 {
		return true
	}

	return false
}

func main() {
	defer funcao1()
	// DEFER = ADIAR
	//É uma declaração que adia a execução de uma função até o final do bloco atual
	funcao2()
	fmt.Println(alunoEstaAprovado(4, 3))
}
