package main

import (
	"fmt"

	"github.com/josebruno2020/algoritmos_go/recursiva"
)

func main() {
	// Busca binaria
	// lista := []int{1, 5, 10, 11, 20}
	// index := busca_binaria.BuscaBinaria(lista, 1)
	// fmt.Printf("Posicao %d da lista \n", index)
	// fmt.Println("Hello world!")

	// Ordenacao por Selecao
	// lista := []int{5, 6, 7, 1, 2, 3, 10, 9}
	// listaOrdenada := ordenacao.OrdenaPorSelecao(lista)
	// fmt.Printf("Antes: %v \n", lista)
	// fmt.Printf("Depois: %v \n", listaOrdenada)

	// Recursiva
	recursiva.Recursiva(5)

	fatorial := recursiva.Fatorial(5)
	fmt.Println("Fatorial de 5 é: ", fatorial)

	lista := []int{2, 4, 6, 5, 8, 2, 500}

	somaArray := recursiva.SomaItemsLista(lista)
	fmt.Println(somaArray)

	itemsNaLista := recursiva.ContaItemsLista(lista)
	fmt.Println("Items na lista: ", itemsNaLista)

	maiorNaLista := recursiva.MaiorItemLista(recursiva.MaiorItemListaParams{Lista: lista})
	fmt.Println("Maior na lista: ", maiorNaLista)
}
