package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(&wg)
	}
	wg.Wait()
	fmt.Println("counter =", counter)
	fmt.Println("Done")
}

func fn(wg *sync.WaitGroup) {
	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()
	wg.Done()
}
