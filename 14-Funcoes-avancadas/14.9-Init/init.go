package main

import "fmt"

var num int

func init() {
	fmt.Println("Função init executada")
	num = 10
}

func main() {
	fmt.Println("Função main executada")
	fmt.Println(num)
}
