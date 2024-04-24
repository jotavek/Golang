package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A ideia principal é pegar dois ou mais canais e juntar em um canal só.

func main() {
	canal := multiplexar(escrever("Bom dia"), escrever("Como vai você?"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// A função multiplexar recebe dois canais como entrada e vai retornar um canal só como saída, todos esses canais só recebem valores
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	// Esse select vai ver qual dos dois canais tem uma mensagem disponivel para ser lida, qual estiver disponivel ele vai jogar no canal de saida, resumindo ele vai encaminhar essas mensagens
	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		// Loop infinito para enviar continuamente valores para o canal
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	// Retorna um canal que espera receber informações.
	return canal
}
