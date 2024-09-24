package busca_binaria

import "fmt"

func BuscaBinaria(lista []int, item int) int {
	min := 0
	max := len(lista) - 1
	i := 0
	for min <= max {
		i++
		meioLista := (min + max) / 2
		chute := lista[meioLista]
		fmt.Printf("tentativa %d, chute: %d. Meio da lista: %d \n", i, chute, meioLista)

		if chute == item {
			return meioLista
		}

		if chute > item {
			max = meioLista - 1
		}

		if chute < item {
			min = meioLista + 1
		}
	}
	return -1
}
