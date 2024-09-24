package ordenacao

import "github.com/josebruno2020/algoritmos_go/helper"

func buscaMenor(lista []int) (int, int) {
	menor := lista[0]
	menorIndice := 0

	for i, element := range lista {
		if element < menor {
			menor = element
			menorIndice = i
		}
	}

	return menorIndice, menor
}

func OrdenaPorSelecao(lista []int) []int {
	novaLista := []int{}
	for range lista {
		menorIndice, menorValor := buscaMenor(lista)
		lista = helper.RemoveIndiceDaLista(menorIndice, lista)
		novaLista = append(novaLista, menorValor)
	}
	return novaLista
}
