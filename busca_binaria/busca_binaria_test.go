package busca_binaria_test

import (
	"testing"

	"github.com/josebruno2020/algoritmos_go/busca_binaria"
)

func TestBuscaBinaria(t *testing.T) {
	lista := []int{1, 2, 4, 6, 8, 10}
	index := busca_binaria.BuscaBinaria(lista, 4)
	expectedResult := 2

	t.Log(index)

	if index != expectedResult {
		t.Fatalf("Resultado esp: %d. Resultado obtido: %d", expectedResult, index)
	}
}
