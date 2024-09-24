package recursiva_test

import (
	"testing"

	"github.com/josebruno2020/algoritmos_go/recursiva"
)

func TestSomaItemsLista(t *testing.T) {
	lista := []int{2, 5, 10, 3}
	somaItems := recursiva.SomaItemsLista(lista)
	expectedResult := 20
	t.Log(somaItems)

	if somaItems != expectedResult {
		t.Fatalf("Resultado esperado: %d. Resultado obtido: %d", expectedResult, somaItems)
	}
}
