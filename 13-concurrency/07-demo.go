package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter = new(Counter)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(&wg)
	}
	wg.Wait()
	fmt.Println("counter =", counter.count)
	fmt.Println("Done")
}

func fn(wg *sync.WaitGroup) {
	counter.Increment()
	wg.Done()
}
