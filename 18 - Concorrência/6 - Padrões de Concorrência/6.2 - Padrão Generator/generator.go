package main

import (
	"fmt"
	"time"
)

// Padrão generator, basicamente esse cara vai ser uma função que vai encapsular uma chamada para uma goroutine e devolver um canal para gente.

func main() {
	canal := escrever("Olá Mundo!")

	// Esse for imprime 5 vezes o que tem no canal
	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
}

// escrever é uma função que retorna um canal de string de uma direção só.
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		// Loop infinito para enviar continuamente valores para o canal
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	// Retorna um canal que espera receber informações.
	return canal
}
