package recursiva

import (
	"github.com/josebruno2020/algoritmos_go/helper"
)

func SomaItemsLista(lista []int) int {
	lenLista := len(lista)
	if lenLista == 0 {
		return 0
	}
	novoArray := helper.RemoveIndiceDaLista(0, lista)
	return lista[0] + SomaItemsLista(novoArray)
}
