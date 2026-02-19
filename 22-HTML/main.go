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

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		usuario := usuario{"João", "joao@go.dev"}
		templates.ExecuteTemplate(w, "home.html", usuario)
	})

	fmt.Println("Servidor rodando na porta 5001S")

	log.Fatal(http.ListenAndServe(":5001", nil))

}
