// Padrão: nome do arquivo = nome do arquivo contendo o programa a ser testado + _test
package media

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

// Padrão: Test + nome da função a ser testada
func TestMedia(t *testing.T) { // O parâmetro t serve para auxiliar
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

/* Execução no terminal:
go test -> executa o programa de teste na pasta atual
go test ./... -> executa todos os programas de teste em todos os diretórios/subdiretórios */
