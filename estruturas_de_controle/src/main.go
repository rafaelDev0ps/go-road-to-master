package main

import "fmt"

func main() {
	var teste = 10
	var code = "@"

	// Estrutura do if
	if teste < 5 && code == "@" {
		fmt.Println("bla bla")
	} else if code != "!" {
		fmt.Println("opa")
	}

	// Estrutura do for
	var cobaia = 0
	for i := 0; i < teste; i++ {
		cobaia = cobaia + i
		fmt.Println("=>", cobaia)
	}
	fmt.Println("Cobaia eh ", cobaia)

	// Estrutura do for range
	var array = [2]string{"ar", "pt"}
	for _, el := range array {
		fmt.Println(el)
	}

	// Estrutura do switch case
	var exp = 2
	switch exp {
	case 10:
		fmt.Println("eh deiz")
	case 5:
		fmt.Println("eh cincu")
	default:
		fmt.Println("eh defo")
	}

}
