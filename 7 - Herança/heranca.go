package main

import "fmt"

type pessoa struct {
	nome      string
	sobreNome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	escola string
	ano    string
}

func main() {
	p1 := pessoa{"Joao", "Victor", 22, 1.63}
	fmt.Println(p1)
	estudante1 := estudante{p1, "IACI", "9 ano"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.altura)

}
