package ordenacao_test

import (
	"testing"

	"github.com/josebruno2020/algoritmos_go/ordenacao"
)

func TestOrdenacaoPorSelecao(t *testing.T) {
	lista := []int{5, 4, 3, 2, 1}
	listaOrdenada := ordenacao.OrdenaPorSelecao(lista)
	expectedResult := []int{1, 2, 3, 4, 5}

	t.Log(listaOrdenada)
	if listaOrdenada[0] != expectedResult[0] || len(listaOrdenada) != len(expectedResult) {
		t.Fatalf("Resultado esp: %v. Resultado obtido: %v", expectedResult, listaOrdenada)
	}
}
