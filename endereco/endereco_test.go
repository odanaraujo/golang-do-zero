package endereco

import "testing"

type cenarioDeTeste struct {
	enderecoInformado string
	enderecoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Agamenon", "Avenida"},
		{"Estrada dos Remédios", "Estrada"},
		{"Rodovia PE 15", "Rodovia"},
		{"Praça das Rosas", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInformado)

		if retornoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado. Tipo recebido: %s | Tipo esperado: %s ",
				retornoRecebido, cenario.enderecoEsperado)
		}
	}
}

//command important - go test --coverprifile result.txt
//command important  - go tool cover --html=result.txt
