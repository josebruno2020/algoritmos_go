package recursiva_test

import (
	"testing"

	"github.com/josebruno2020/algoritmos_go/recursiva"
)

func TestFatorialFunction(t *testing.T) {
	numToTest := 5
	result := recursiva.Fatorial(numToTest)
	expetedResult := 120
	t.Log(result)
	if result != expetedResult {
		t.Fatalf("Resultado esperado: %d. Esperado recebido: %d", expetedResult, result)
	}
}
