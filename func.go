package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}
}