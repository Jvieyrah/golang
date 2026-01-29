package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1) // chamada recursiva onde a função chama ela mesma para obter o resultado
}

func main() {
	fmt.Println(fibonacci(10))
}
