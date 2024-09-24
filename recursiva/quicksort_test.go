package recursiva_test

import (
	"testing"

	"github.com/josebruno2020/algoritmos_go/recursiva"
)

func TestQuicksort(t *testing.T) {
	lista := []int{2, 4, 1, 5, 3}
	listaOrdenada := recursiva.Quicksort(lista)
	expectedResult := []int{1, 2, 3, 4, 5}

	t.Logf("Lista ordenada: %v\n", listaOrdenada)
	for i, element := range listaOrdenada {
		if element != expectedResult[i] {
			t.Fatalf("Lista %v não está compativel na posicao %d", listaOrdenada, i)
		}
	}
}
