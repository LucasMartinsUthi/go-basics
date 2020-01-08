package main

import (
	"fmt"
	"os"
	"strconv"
)

func print (text string, value int){
	fmt.Printf("%s %d\n", text, value)
}

func soma (a, b int) int {
	return a + b
}

func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("Uso: soma <valores>")
		os.Exit(1)
	}

	var total int

	for _, v := range os.Args[1 : len(os.Args)] {
		v, _ := strconv.Atoi(v)

		total = soma(total, v)
	}

	fmt.Printf("Soma total: %d\n", total)
} 	