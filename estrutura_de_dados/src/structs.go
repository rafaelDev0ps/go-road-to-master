package main

import "fmt"

func main() {

	type comida struct {
		validade  string
		valor     float32
		descricao string
	}

	alimento := &comida{
		validade:  "12/12/2020",
		valor:     25.99,
		descricao: "X-Burger",
	}

	// Acessando a struct
	fmt.Println("Alimento => ", alimento)
	fmt.Println("Alimento => ", *alimento) // quando se tem a posicao na memoria (&)

	// Acessando elemento da struct
	fmt.Println("Descr. alimento => ", alimento.descricao)

	// Acessando o ponteiro que indica a posicao na memoria da struct
	fmt.Println("Mem. struct => ", &alimento)

}
