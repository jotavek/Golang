package main

import (
	"fmt"
	"time"
)

// A goroutine é uma função ou um método que é invocado dessa forma
// go escrever("Texto")
// Ela executa mas nao espera ela terminar de executar para seguir o programa
func main() {
	// CONCORRENCIA != PARALELISMO
	go escrever("Olá Mundo!")
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
