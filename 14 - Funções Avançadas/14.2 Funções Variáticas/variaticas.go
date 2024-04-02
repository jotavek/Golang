package main

import "fmt"

//Voce só pode ter um parametro variatico por função "numeros ...int" por exemplo e ele deverá ser o ultimo parametro colocado na função
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//func main() {
//	totaldaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9)
//	fmt.Println(totaldaSoma)
//}


func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma (1, 2, 3, 0)
	fmt.Println(totalDaSoma)
	escrever ("Olá mundo", 10, 4, 5, 6, 7)
}