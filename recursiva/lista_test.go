package recursiva_test

import (
	"fmt"
	"testing"

	"github.com/josebruno2020/algoritmos_go/recursiva"
)

var lista = []int{1, 5, 2, 10}

func TestContaItemNaLista(t *testing.T) {
	totalItems := recursiva.ContaItemsLista(lista)
	expectedResult := 4
	fmt.Println(totalItems)
	if totalItems != expectedResult {
		t.Fatalf("Resultado esperado: %d. Resultado obtido: %d", expectedResult, totalItems)
	}
}

func TestMaiorItemNaLista(t *testing.T) {
	maiorItem := recursiva.MaiorItemLista(recursiva.MaiorItemListaParams{Lista: lista})
	expectedResult := 10

	t.Log(maiorItem)
	if maiorItem != expectedResult {
		t.Fatalf("Resultado esperado: %d. Resultado obtido: %d", expectedResult, maiorItem)
	}
}
