package main

import (
	"fmt"
	"time"
)

func main() {

	canal := escrever("Olá Mundo!")
	for valor := range canal {
		fmt.Println(valor)
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal

}
