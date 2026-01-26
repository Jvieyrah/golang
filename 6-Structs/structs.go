package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	end   endereco
}

type endereco struct {
	logradouro string
	numero     uint8
	cep        string
}

func main() {
	var u usuario
	u.nome = "João"
	u.idade = 20
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua 1", 123, "12345-678"}
	u2 := usuario{"João", 20, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Pedro"}
	fmt.Println(u3)

	u4 := usuario{idade: 100}
	fmt.Println(u4)
}
