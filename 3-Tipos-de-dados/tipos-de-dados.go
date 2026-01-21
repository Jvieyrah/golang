package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero1 int8 = 100
	var numero2 int16 = 20000
	var numero3 int32 = 300000000
	var numero4 int64 = 4000000000000000000

	var numero int = 1111111111111111111 // 64 bits

	numeroInferido := -1000000000000000000 // 64 bits

	fmt.Println(numero1, numero2, numero3, numero4)
	fmt.Println(numero)
	fmt.Println(numeroInferido)

	var numero5 uint8 = 100
	var numero6 uint16 = 20000
	var numero7 uint32 = 300000000
	var numero8 uint64 = 4000000000000000000

	fmt.Println(numero5, numero6, numero7, numero8)

	//alias
	var numero9 rune = 12345 // int32 = 4 bytes = carateres da tabela ascii.
	var numero10 byte = 123  // int8 = 1 byte = 0 a 255

	fmt.Println(numero9, numero10)

	//numeros reais com ponto ao inves de virgula.
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 123456.789
	numerotal3 := 123456.789

	fmt.Println(numeroReal1, numeroReal2)
	fmt.Println(numerotal3)

	//string
	var str string = "texto"
	fmt.Println(str)

	str2 := "texto 2"
	fmt.Println(str2)

	char := 'b'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	// nulos

	var inteiro int
	fmt.Println(inteiro)

	var booleano bool
	fmt.Println(booleano)

	var erro error
	fmt.Println(erro)

	//errors
	var erro2 error = errors.New("erro interno")
	fmt.Println(erro2)

}
