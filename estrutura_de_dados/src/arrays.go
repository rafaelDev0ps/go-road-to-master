package main

import "fmt"

func main() {
	var vetor []int

	// array com tamanho fixo
	var vetor2 = [3]int{12, 45, 76}

	// declarando array utilizando variadic
	var vetor3 = [...]int{10, 9, 8, 4, 5}

	fmt.Println("vetor2 =>", vetor2[0])
	fmt.Println("teste github workspace")

	for i, el := range vetor3 {
		fmt.Print("indice =>", i)
		fmt.Println(" | elemento =>", el)
	}

	// Inserindo elemento em um vetor
	vetor = append(vetor, 10)
	fmt.Println("Inserindo no vetor =>", vetor)

}
