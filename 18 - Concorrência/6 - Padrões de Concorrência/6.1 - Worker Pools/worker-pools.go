package main

import "fmt"

// O que são Worker Pools?
// É como se tivesse  uma fila de tarefas para serem executadas e voce tem varios workers, varios processos pegando itens dessa fila de maneira independente

func main() {
	tarefas := make(chan int, 45)    // Esse canal vai ter a sequencia de numeros que nós precisamos calcular
	resultados := make(chan int, 45) // Esse canal vai armazenar o resultado

	go worker(tarefas, resultados)

	// Esse for manda numeros para o canal tarefas
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	// For para printar os resultados
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

//Eu posso especificar como parâmetro qual canal vai receber informações e qual canal vai enviar informações
// tarefas <-chan int é um canal que só vai receber dados
// tarefas chan<- int é um canal que só vai enviar dados
// A função worker vai esperar receber uma tarefa do for, enquanto ele recebe a tarefa ele vai executando a função de fibonacci e vai mandando pro canal de resultados.
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
