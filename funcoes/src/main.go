package main

import "fmt"

// exemplo de uma funcao com parametros
func show(nome string) (string, error) {

	if nome == "" {
		return "", fmt.Errorf("falta o argumento nome")
	}
	return nome, nil
}

// Exemplo de funcao com defer
func chave() int {
	var teste = 4

	defer func() { // defer ignora valores de retorno!
		teste = teste + 5
		fmt.Println("parte defer =>", teste)
	}()

	teste = teste * teste
	fmt.Println("parte chave =>", teste)
	return teste
}

func main() {
	titulo, err := show("")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(titulo)

	ex := chave()

	fmt.Println("parte ex =>", ex)
}
