package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	contador := 0

	var wg sync.WaitGroup
	var mute sync.Mutex

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mute.Lock()
			aux := contador
			runtime.Gosched()
			aux++
			contador = aux
			mute.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutine", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("O contador eh", contador)
}
