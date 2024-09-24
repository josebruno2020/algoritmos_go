package recursiva

import "fmt"

func getMenoresEMaiores(lista []int, pivo int) ([]int, []int) {
	var menores, maiores []int

	for _, element := range lista[1:] {
		if element <= pivo {
			menores = append(menores, element)
		} else {
			maiores = append(maiores, element)
		}
	}

	return menores, maiores
}

func Quicksort(lista []int) []int {
	if len(lista) < 2 {
		return lista
	}
	pivo := lista[0]
	menores, maiores := getMenoresEMaiores(lista, pivo)
	fmt.Println("pivo: ", pivo)

	return append(append(Quicksort(menores), pivo), Quicksort(maiores)...)
}
