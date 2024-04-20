package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Mensagem do CANAL 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Mensagem do CANAL 2"
		}
	}()
	// O loop infinito abaixo mantém o programa em execução, já que as goroutines estão enviando mensagens continuamente para os canais.
	for {
		// O bloco select aguarda a disponibilidade de dados nos canais canal1 e canal2 e imprime as mensagens assim que elas estão prontas para serem lidas.
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
