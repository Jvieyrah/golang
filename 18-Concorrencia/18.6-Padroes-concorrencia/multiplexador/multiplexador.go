package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canalDeEntrada1 := escrever("Ola Mundo")
	canalDeEntrada2 := escrever("Programando em Go")

	canalDeSaida := multiplexar(canalDeEntrada1, canalDeEntrada2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-canalDeSaida)
	}

}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem1 := <-canalDeEntrada1:
				canalDeSaida <- mensagem1
			case mensagem2 := <-canalDeEntrada2:
				canalDeSaida <- mensagem2
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}
