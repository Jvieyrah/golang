package main

import (
	"fmt"
)

func main() {

	// var i int = 0
	// for i < 10 {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("O valor de i é", i)
	// 	i++
	// }

	// for j := 0; j < 10; j++ {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("O valor de j é", j)
	// }

	nomes := []string{"João", "Maria", "Pedro"}
	for _, nome := range nomes {
		fmt.Println("O nome é", nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuarios := map[string]string{
		"nome":      "joao",
		"sobrenome": "silva",
		"profissao": "dev",
		"cidade":    "Recife",
	}
	for chave, valor := range usuarios {
		fmt.Println(chave, valor)
	}

}
