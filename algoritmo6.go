package main

import (
	"fmt"
)

func main() {
	// Criando o array de inteiros
	array := [10]int{10, 25, 36, 42, 55, 68, 73, 81, 94, 100}

	var valor int
	fmt.Print("Digite um valor para buscar no array: ")
	fmt.Scan(&valor)

	encontrado := false
	for _, v := range array {
		if v == valor {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println("O valor", valor, "está presente no array.")
	} else {
		fmt.Println("O valor", valor, "não está presente no array.")
	}
}
