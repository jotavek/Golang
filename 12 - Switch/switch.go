package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	default:
		return "Numero inválido"
	}
}
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	default:
		return "Numero inválido"
	}
}

func main() {
	dia := diaDaSemana(4)
	fmt.Println(dia)
}