package main

import (
	"fmt"
	"sync"
)

func jiva(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go jiva(&w, ch)

	}
	w.Wait()
	fmt.Println("final value of x:", x)
}
