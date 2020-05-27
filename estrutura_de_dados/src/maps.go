package main

import "fmt"

func main() {
	var map1 = map[string]int{"nos": 4, "vos": 5}

	map1["eu"] = 1
	map1["vc"] = 2
	map1["eles"] = 3

	// Acessando o map
	fmt.Println("Map => ", map1)

	// Acessando um valor especifico pelo rotulo
	fmt.Println("Valor esp => ", map1["eu"])

	fmt.Println("Buscando o abacate => ", map1["abacate"])

	map2 := map[string]string{}

	// Se o valor que buscamos pela chave nao existe, o Go retorna um valor-zero
	fmt.Println("Buscando o abacate => ", map2["abacate"])

	// Removendo um elemento do map
	delete(map1, "eu")
	fmt.Println("Map1 => ", map1)

	// Acessando a chave e campo de existencia do map
	valor, exists := map1["eles"]
	fmt.Printf("Valor %d | Existe %t\n", valor, exists)

	// Percorrendo um map
	for chave, valor := range map1 {
		fmt.Printf("Chave %s | Valor %d\n", chave, valor)
	}

}
