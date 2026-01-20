package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Arquivo Main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("joao@go.dev")
	fmt.Println(erro)
}
