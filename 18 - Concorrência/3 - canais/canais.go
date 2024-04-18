package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	// Coloca o que o canal vai receber em uma variável e printa essa variavel pra ver o resultado
	mensagem := <-canal
	fmt.Println(mensagem)

	for msg := range canal {
		fmt.Println(msg)
	}
	fmt.Println("Fim do programa.")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	//Essa função fecha o canal quando acabar o looping, ou seja ele nao vai enviar nem receber mais dados. Evitando o DEADLOCK.
	close(canal)
}
