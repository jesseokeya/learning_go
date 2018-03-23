package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter = 0

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func increment() {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}

func sayHello() {
	m.Lock()
	fmt.Printf("Hello World #%v\n", counter)
	m.Unlock()
	wg.Done()
}
