package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}

func main() {

	canal1 := make(chan string)

	go escrever("Canal 1 ", canal1)
	fmt.Println("Depois da funcao escrever")

	// for {
	// 	mensagem, aberto := <-canal1

	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal1 {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim")

}
