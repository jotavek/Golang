package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	// Primeira maneira declarando cada variavel da estrutura
	var u usuario
	u.nome = "João"
	u.idade = 20
	fmt.Println(u)

	// Segunda forma com inferencia de tipo
	enderecoExemplo := endereco{"Rua dos bobos", 0}
	u2 := usuario{"João Victor", 21, enderecoExemplo}
	fmt.Println(u2)

	// Terceira forma declarando apenas uma variavel da estrutura
	usuario3 := usuario{idade: 21}
	fmt.Println(usuario3)
}
