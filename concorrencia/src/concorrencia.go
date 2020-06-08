package main

import (
	"fmt"
	"sync"
)

// Add, Done, Wait

var waitingGroup sync.WaitGroup

func main() {

	waitingGroup.Add(2)

	go funcaoUm()
	go funcaoDois()

	waitingGroup.Wait()
}

func funcaoUm() {
	fmt.Println("Funcao 1")
	arrayUm := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		fmt.Println(arrayUm[i])
	}
	// Done
	waitingGroup.Done()
}

func funcaoDois() {
	fmt.Println("Funcao 2")
	arrayDois := [10]int{12, 43, 65, 89, 9, 74, 19, 15, 34, 64}
	for i := 0; i < 10; i++ {
		fmt.Println(arrayDois[i])
	}

	waitingGroup.Done()
}
