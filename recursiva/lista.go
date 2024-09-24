package recursiva

import (
	"github.com/josebruno2020/algoritmos_go/helper"
)

type MaiorItemListaParams struct {
	Lista []int
	Valor int
}

// Conta quantos elementos uma lista possui, de forma recursiva
func ContaItemsLista(lista []int) int {
	if len(lista) == 0 {
		return 0
	}

	novoArray := helper.RemoveIndiceDaLista(0, lista)
	return 1 + ContaItemsLista(novoArray)
}

// Encontrar o maior elemento em uma lista, de forma recursiva
func MaiorItemLista(params MaiorItemListaParams) int {
	// fmt.Println("Valor: ", params.Valor)
	// fmt.Println("Lista: ", params.Lista)
	if len(params.Lista) == 0 {
		return 0
	}
	if len(params.Lista) == 1 {
		return params.Valor
	}
	novaLista := helper.RemoveIndiceDaLista(0, params.Lista)
	novoValor := novaLista[0]

	if novoValor > params.Valor {
		return MaiorItemLista(MaiorItemListaParams{Lista: novaLista, Valor: novoValor})
	}

	return MaiorItemLista(MaiorItemListaParams{Lista: novaLista, Valor: params.Valor})
}
