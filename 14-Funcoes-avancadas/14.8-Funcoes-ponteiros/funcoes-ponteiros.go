package main

import "fmt"

func inverterSinalDoNumero(num int) {
	num = num * -1
	fmt.Println(num)
}

func inverterSinalDoNumeroPonteiro(num *int) {
	*num = *num * -1
	fmt.Println(*num)
}

func main() {
	numero := 20
	inverterSinalDoNumero(numero)
	fmt.Println(numero)
	inverterSinalDoNumeroPonteiro(&numero)
	fmt.Println(numero)

	novoNumero := 30
	inverterSinalDoNumeroPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
