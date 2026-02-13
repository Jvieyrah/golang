package main

import (
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
	usuarioEmJson := `{"nome":"João","idade":20,"cursos":[{"nome":"Golang","categoria":"Programação"},{"nome":"Python","categoria":"Programação"}]}`

	var u usuario

	if error := json.Unmarshal([]byte(usuarioEmJson), &u); error != nil {
		log.Fatal(error)
	}

	fmt.Println(u)

	usuarioEmJson2 := `{"nome":"João","idade":20,"cursos":[{"nome":"Golang","categoria":"Programação"},{"nome":"Python","categoria":"Programação"}]}`

	var u2 map[string]interface{}

	if error := json.Unmarshal([]byte(usuarioEmJson2), &u2); error != nil {
		log.Fatal(error)
	}

	fmt.Println(u2)

}
