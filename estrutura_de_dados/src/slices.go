package main

import "fmt"

// Leia: https://blog.golang.org/slices-intro

func main() {
	// um slice funciona "em cima" de um array

	var slice = make([]float64, 3) // mais performatico
	var slice2 = []float64{}

	// Inserindo elementos em um slice
	slice = append(slice, 16.04)
	slice = append(slice, 34.12)
	slice = append(slice, 19.96)

	slice2 = append(slice2, 16.04)
	slice2 = append(slice2, 34.12)
	slice2 = append(slice2, 19.96)

	fmt.Println("slice com make => ", slice)
	// Output: slice =>  [0 0 0 16.04 34.12 19.96]

	// Removendo elementos de um array
	// slice = slice[:len(slice)-1]

	fmt.Println("slice2 => ", slice2)

	slice2 = slice2[2:]

	fmt.Println("slice2 => ", slice2)

	for _, el := range slice2 {
		fmt.Println("Elementos do slice2 => ", el)
	}

}
