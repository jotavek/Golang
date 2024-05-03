package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	// Parse todos os templates no diretório atual com extensão .html e armazena-os em "templates"
	templates = template.Must(template.ParseGlob("*.html"))

	// Manipulador para a rota "/home"
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Maria", "joao.pedro@gmail.com"}
		// Renderiza o template "home.html" usando os dados do usuário "u" e escreve a saída para o ResponseWriter "w"
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
