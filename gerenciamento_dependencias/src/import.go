package main

import (
	interfaces "aprendendo/interfaces/src"
	random "aprendendo/random/src"
	"fmt"
)

type generatorHandler struct {
	random.Generator
}

func (gh generatorHandler) numberGenerator(ng random.NumberGenerator) int64 {
	return ng.GerarNumero(gh.Seed)
}

func main() {
	autor := interfaces.Autor{
		Nome:  "Rafael Hurrycaner",
		Idade: 25,
	}

	maquina := interfaces.Computador{
		Modelo: "XZ064",
		Marca:  "Acer",
	}

	digitar := interfaces.Digitar(maquina)
	pintor := canetinha(autor)
	fmt.Println(digitar)
	fmt.Println(pintor)

	// Gerando numeros aleatorios

	generator := generatorHandler{
		random.Generator{
			Seed: 10,
		},
	}

	randomNum := generator.numberGenerator(generator)
	fmt.Println(randomNum)
}
func canetinha(l interfaces.Livro) string {
	return l.Escreve()
}
