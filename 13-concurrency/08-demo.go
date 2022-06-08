package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

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
	atomic.AddInt64(&counter, 1)
	wg.Done()
}
