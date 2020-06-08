package main

import "fmt"

func main() {

	val, err := cobaia(16)
	if err != nil {
		panic(err) // termina o programa durante execucao
	} else {
		fmt.Println(val)
	}

	fmt.Println("Depois de tudo...")
}

func cobaia(index int) (int, error) { // funcao retorna erro, caso tenha
	teste := [5]int{1, 2, 3, 4, 5}

	if index > len(teste) {
		return 0, fmt.Errorf("Valor de indice invalido") // retornando erro
	}

	return teste[index], nil
}
