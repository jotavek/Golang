package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE
	// CREATE - POST
	// READ - GET
	// UPDATE - PUT
	// DELETE - DELETE
	// Ele vai conter todas as rotas que nós vamos utilizar para fazer as interações com o banco de dados
	router := mux.NewRouter()
	// Quando a minha rota de /usuarios receber um POST ela vai executar a função CriarUsuario que está na pasta servidor.
	// Eu posso ter mais de uma rota com o mesmo URI (usuarios) desde que o método seja diferente.
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
