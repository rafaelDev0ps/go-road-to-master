package main

import "fmt"

func main() {
	type stream struct {
		viewers int
		canal   string
	}

	plataforma := map[int]*stream{}

	// Podemos atribuir qualquer tipo de variavel em um map em Go
	plataforma[1] = &stream{
		viewers: 9345098234,
		canal:   "rafaLiveCoding",
	}

	fmt.Println("Map com struct => ", plataforma[1])
}
