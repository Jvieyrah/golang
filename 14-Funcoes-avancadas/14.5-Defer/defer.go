package main

import "fmt"

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será:")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Defer")
	defer fmt.Println("Primeiro")
	fmt.Println("Segundo")
	fmt.Println(alunoEstaAprovado(8, 6))
}
