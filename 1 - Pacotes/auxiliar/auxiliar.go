package auxiliar

import "fmt"

//Essa função escreve na sua tela e chama a função escrever2 que está no mesmo package auxiliar
func Escrever() {
	fmt.Println("Escrevendo da pasta auxiliar")
	escrever2()
}
