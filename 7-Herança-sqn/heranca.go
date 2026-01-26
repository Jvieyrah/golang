package main

import "fmt"

func main() {
	p1 := pessoa{"João", 20, "M"}
	estudante1 := estudante{p1, "Engenharia", "C", "USP"}
	fmt.Println(estudante1)

	fmt.Println(estudante1.nome)
	fmt.Println(estudante1.curso)
}

type pessoa struct {
	nome  string
	idade uint8
	sexo  string
}

type estudante struct {
	pessoa
	curso     string
	turma     string
	faculdade string
}
