package main

import "fmt"

func generica(a interface{}) {
	fmt.Println(a)
}

func main() {
	generica(1)
	generica("String")
	generica(true)

	mapa := map[interface{}]interface{}{
		"nome":         "joao",
		"idade":        20,
		1:              true,
		float32(29.89): func() {},
	}
	for chave, valor := range mapa {
		fmt.Println(chave, valor)

	}
}
