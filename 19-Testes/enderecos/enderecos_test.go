package enderecos_test

import (
	. "introducao-teste/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido  string
	resultadoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	cenarioDeTestes := []cenarioDeTeste{
		{"Rua Paulista", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Paulista", "Rodovia"},
		{"Praça Paulista", "tipo inválido"},
		{"", "tipo inválido"},
		{"RUA PAULISTA", "Rua"},
		{"AVENIDA PAULISTA", "Avenida"},
		{"RODOVIA PAULISTA", "Rodovia"},
		{"PRAÇA PAULISTA", "tipo inválido"},
	}

	for _, cenario := range cenarioDeTestes {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if retornoRecebido != cenario.resultadoEsperado {
			t.Errorf("Esperado '%s', mas obteve '%s'", cenario.resultadoEsperado, retornoRecebido)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("1 é maior que 2, Quebrou o teste")
	}
}
