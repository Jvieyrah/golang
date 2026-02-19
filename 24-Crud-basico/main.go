package main

import (
	"crud-basico/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuarios).Methods(http.MethodPost)

	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)

	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)

	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)

	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Servidor rodando na porta 5005")
	log.Fatal(http.ListenAndServe(":5005", router))
}
