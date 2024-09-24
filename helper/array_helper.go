package helper

func RemoveIndiceDaLista(indice int, lista []int) []int {
	novaLista := []int{}
	for i, element := range lista {
		if i != indice {
			novaLista = append(novaLista, element)
		}
	}
	return novaLista
}
