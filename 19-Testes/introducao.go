package main

import (
	"fmt"
	"introducao-teste/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Avenida Paulista")

	fmt.Println(tipoDeEndereco)
}
