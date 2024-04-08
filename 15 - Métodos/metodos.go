package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// É como se fosse uma função que tem o nome salvar(), porém ela está associada a uma certa struct ou método, que nesse caso é uma struct usuario
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() {
	if u.idade > 18 {
		fmt.Printf("Usuário %s é maior de idade\n", u.nome)
	} else {
		fmt.Printf("Usuário %s é menor de idade\n", u.nome)
	}
}
func (u usuario) maiorDeIdadeTouF() bool {
	return u.idade >= 18
}

// Esse método altera um valor dentro do usuario, nesse caso usamos o ponteiro, acrescentando * no usuario, neste caso nao precisa fazer a desreferenciação com o &
func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	apelido := usuario{"Joao", 22}
	fmt.Println(apelido)
	apelido.salvar()

	usuario2 := usuario{"Diego", 45}
	usuario2.aniversario()
	usuario2.salvar()
	usuario2.maiorDeIdade()
	maiorDeIdade := usuario2.maiorDeIdadeTouF()
	fmt.Println(maiorDeIdade)
	fmt.Println(usuario2)

}
