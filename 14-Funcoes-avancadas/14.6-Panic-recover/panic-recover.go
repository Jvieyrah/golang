package main

import "fmt"

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperaExecucao()
	defer fmt.Println("Média calculada. Resultado será: ")
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	}
	if media < 6 {
		return true
	}
	panic("A média é exatamente 6")
}

func recuperaExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 7))
	fmt.Println(alunoEstaAprovado(6, 6))
}
