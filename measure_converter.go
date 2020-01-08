package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args array de entrada
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	// unidade := map[string]string{
	// 	"origem" : "",
	// 	"destino" : "",
	// }
	
	// valor := map[string][]float32{
	// 	"origem" : []float32{},
	// 	"destino" : []float32{},
	// }

	unidadeOrigem := os.Args[len(os.Args)-1]
	valorOrigem := os.Args[1 : len(os.Args)-1]

	var valorDestino float64
	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"

	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"

	} else {
		fmt.Println("%s Unidade não reconhecida", unidadeOrigem)
		os.Exit(1)
	}

	for i, v := range valorOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf(
				"valor %s na posição %d não é um número válido!",
				v, i)

			os.Exit(1)
		}

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf(
			// resultado
			"%.2f %s = %2.f %s\n",
			valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}
}