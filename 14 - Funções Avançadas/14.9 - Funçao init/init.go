package main

import "fmt"

//Função init inicia primeiro no arquivo, eu posso ter uma em cada arquivo, e nao por pacote.
// É bom para fazer um setup, inicializar alguma variavel que vai ser utilizada..
func main() {
	fmt.Println("Print da função main")

}

func init() {
	fmt.Println("Print da função init")
}