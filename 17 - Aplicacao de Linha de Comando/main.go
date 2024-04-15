package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

// Para executar a aplicação é preciso usar o parâmetro .run() passando "os.Args" como padrão, assim os comandos do sistema operacional são reconhecidos pela linha de comando
// Observação, o método .run retorna um erro, entao é preciso tratar esse erro caso ele ocorra
//Se ocorrer um erro o programa para na hora
func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
