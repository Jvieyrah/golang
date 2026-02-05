package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetrasMinusculas := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoEmLetrasMinusculas, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {

		return strings.Title(primeiraPalavra)
	}

	return "tipo inválido"
}
