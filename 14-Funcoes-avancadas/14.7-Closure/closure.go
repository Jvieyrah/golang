package main

import "fmt"

func closure() func() {
	texto := "Texto da função"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

}

func main() {
	texto := "Texto da main"
	fmt.Println(texto)
	funcao := closure()
	funcao()
}
