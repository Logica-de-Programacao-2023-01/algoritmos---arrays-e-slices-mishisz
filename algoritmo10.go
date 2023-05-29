package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 25, 36, 42, 55, 68, 73, 81, 94, 100}

	min := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	fmt.Println("Valor mínimo:", min)
	fmt.Println("Valor máximo:", max)
}
