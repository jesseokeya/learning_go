package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func(ch <-chan int) {
			i := <-ch
			m.Lock()
			fmt.Println(i)
			m.Unlock()
			wg.Done()
		}(ch)

		go func(ch chan<- int, index int) {
			m.Lock()
			ch <- 3 * index
			m.Unlock()
			wg.Done()
		}(ch, j)
	}
	wg.Wait()
}
