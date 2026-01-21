package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"
	var variavel3 string
	fmt.Println(variavel1, variavel2, variavel3)

	var (
		variavel4 string = "variavel 4"
		variavel5 string = "variavel 5"
	)
	fmt.Println(variavel4, variavel5)

	variavel6, variavel7 := "variavel 6", "variavel 7"
	fmt.Println(variavel6, variavel7)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel6, variavel7 = variavel7, variavel6
	fmt.Println(variavel6, variavel7)
}
