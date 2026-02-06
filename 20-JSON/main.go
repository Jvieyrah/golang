package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type usuario struct {
	Nome   string  `json:"nome"`
	Idade  uint8   `json:"idade"`
	Cursos []curso `json:"cursos"`
}

type curso struct {
	Nome      string `json:"nome"`
	Categoria string `json:"categoria"`
}

func main() {

	u := usuario{
		Nome:  "João",
		Idade: 20,
		Cursos: []curso{
			{
				Nome:      "Golang",
				Categoria: "Programação",
			},
			{
				Nome:      "Python",
				Categoria: "Programação",
			},
		},
	}

	fmt.Println(u)

	studentJson, error := json.Marshal(u)

	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(studentJson)
	fmt.Println(bytes.NewBuffer(studentJson))
	fmt.Println(string(studentJson))

	uList := map[string]string{
		"nome":  "João",
		"nome1": "João",
	}

	uListJson, error := json.Marshal(uList)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(uListJson)
	fmt.Println(bytes.NewBuffer(uListJson))
	fmt.Println(string(uListJson))
}
