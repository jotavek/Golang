package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}


	if outroNumero := 24; outroNumero > 30 {
		fmt.Println("Numero é maior que 30")
	} else if outroNumero < 25{
		fmt.Println("Numero é menor que 25")
	} else {
		fmt.Println("Esse numero está entre 25 e 30")
	}
}