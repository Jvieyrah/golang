package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é impar")
	}

	if 8%4 == 0 {
		fmt.Println("8 é divisível por 4")
	} else if 8%2 == 0 {
		fmt.Println("8 é par")
	} else {
		fmt.Println("8 é impar")
	}

	numero := 9
	if outroNumero := numero; outroNumero%2 == 0 {
		fmt.Println("outroNumero é par")
	} else {
		fmt.Println("outroNumero é impar")
	}

}
