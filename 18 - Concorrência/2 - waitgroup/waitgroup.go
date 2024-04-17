package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup sincroniza duas ou mais goroutines.
// Para criar um waitgroup voce coloca numa variável com sync.WaitGroup e depois de declarar voce coloca waitGroup.add(quantidade de goroutines que irao estar rodando ao mesmo tempo)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1 na contagem
	}()
	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	// Espera a contagem da waitgroup chegar em 0
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
