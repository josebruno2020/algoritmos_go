package helper_test

import (
	"testing"

	"github.com/josebruno2020/algoritmos_go/helper"
)

func TestRemoveIndiceDaLista(t *testing.T) {
	lista := []int{1, 2, 3}
	listLen := len(lista)
	novaLista := helper.RemoveIndiceDaLista(0, lista)
	novaListaLen := len(novaLista)
	t.Log(novaLista)

	if novaLista[0] != 2 {
		t.Fatalf("Nao foi possivel remover a posicao 0")
	}

	if novaListaLen != listLen-1 {
		t.Fatalf("Tamanho esperado: %d. Tamanho da lista: %d", listLen-1, novaListaLen)
	}
}
