package main

import "fmt"

func somar(nr1 int8, nr2 int8) int8 {
	return nr1 + nr2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	resultado := somar(10, 20)
	fmt.Println(resultado)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado2 := f("Texto da função")
	fmt.Println(resultado2)

	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)

	soma2, _ := calculosMatematicos(10, 20) // ignora o valor da subtracao
	fmt.Println(soma2)

}
