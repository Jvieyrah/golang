package main

import "fmt"

func main() {
	retorno := func(txt string) string {
		return fmt.Sprintf("Função Anônima -> %s", txt)
	}("passando parâmetro")

	fmt.Println(retorno)
}
