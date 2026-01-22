package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	var variavel1 string = "Ola"
	variavel2 := " Mundo"
	soma3 := variavel1 + variavel2
	fmt.Println(soma3)

	fmt.Println(1 > 2)  //false
	fmt.Println(1 < 2)  //true
	fmt.Println(1 >= 2) //false
	fmt.Println(1 <= 2) //true

	fmt.Println(1 == 2) //false
	fmt.Println(1 != 2) //true

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro, falso)

	fmt.Println(verdadeiro && verdadeiro) //verdadeiro
	fmt.Println(verdadeiro && falso)      //falso
	fmt.Println(verdadeiro || verdadeiro) //verdadeiro
	fmt.Println(verdadeiro || falso)      //verdadeiro
	fmt.Println(!verdadeiro)              //falso
	fmt.Println(!falso)                   //verdadeiro

	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 10
	fmt.Println(numero)

	numero *= 2
	numero /= 2
	fmt.Println(numero)

	numero %= 2
	fmt.Println(numero)

	// ternarios nao pode ser usados em GO
	// texto := numero > 5  ? "Maior que 5" : "Menor que 5"
	// fmt.Println(texto) < -- ./operadores.go:59:23: invalid character U+003F '?'
	//./operadores.go:59:25: syntax error: unexpected literal "Maior que 5" at end of statement

	//jeito correto
	if numero > 5 {
		fmt.Println("Maior que 5")
	} else {
		fmt.Println("Menor que 5")
	}
}
