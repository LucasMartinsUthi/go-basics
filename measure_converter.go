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

	if unidadeOrigem == "celsisu" {
		unidadeDestino := "fahrenheit" 
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino := "milhas"
	} else {
		fmt.Println("%s Unidade não reconhecida", unidadeOrigem)
		os.Exit(1)
	}
}