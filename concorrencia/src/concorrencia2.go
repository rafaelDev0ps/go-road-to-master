package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go enviadora(canal)
	recebedora(canal)
}

func recebedora(c chan<- int) {
	c <- 123
}

func enviadora(c <-chan int) {
	fmt.Println("O valor do canal eh ", <-c)
}
