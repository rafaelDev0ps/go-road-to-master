package main

import (
	"fmt"
	"reflect"
)

func main() {
	var teste = "string"
	fmt.Println(teste)
	fmt.Println(reflect.TypeOf(teste))

	var teste1 = false
	fmt.Println(teste1)
	fmt.Println(reflect.TypeOf(teste1))

	var teste2 = 54
	fmt.Println(teste2)
	fmt.Println(reflect.TypeOf(teste2))

	var teste3 = 23.3234
	fmt.Println(teste3)
	fmt.Println(reflect.TypeOf(teste3))

}
