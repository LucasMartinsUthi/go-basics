package main

import (
	"fmt"
	"os"
	"strconv"
)

func particionar (numeros []int, pivo int) (menores []int, maiores []int) {
	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}

func quicksort(numeros []int ) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)

	indicePivo := len(n) / 2
	pivo := n[indicePivo]

	n = append(n[:indicePivo], n[indicePivo + 1:]...)

	menores, maiores := particionar(n, pivo)

	return append(
		append(quicksort(menores), pivo), 
		quicksort(maiores)...)

	return numeros
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: quicksort <valores>")
		os.Exit(1)
	}

	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {
		numero, err := strconv.Atoi(n)

		if err != nil {
			fmt.Printf("%s não é um numero definido", n)
		}

		numeros[i] = numero
	}

	fmt.Println(quicksort(numeros))
} 	