package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var resultCh chan int = make(chan int)
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, resultCh)
	wg.Wait()
	result := <-resultCh
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	fmt.Println("	add operation - started")
	time.Sleep(5 * time.Second)
	result := x + y
	ch <- result
	fmt.Println("	add operation - completed")
	wg.Done()
}
