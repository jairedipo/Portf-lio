package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // Permite executar arquivos de teste em paralelo, poupando tempo

	if runtime.GOARCH == "amd64" { // Selecionar uma arquitetura específica
		// Pular um teste. A mesnsagem é logada se executado com 'go test -v'
		t.Skip("Não funciona em arquitetura amd64")
	}
	// ...
	t.Fail()
}
