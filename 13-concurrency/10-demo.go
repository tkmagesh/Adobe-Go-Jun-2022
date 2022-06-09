package main

import (
	"fmt"
	"sync"
	"time"
)

//communicating by sharing memory (DON'T)
var result int

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200)
	wg.Wait()
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int) {
	fmt.Println("	add operation - started")
	time.Sleep(5 * time.Second)
	result = x + y
	fmt.Println("	add operation - completed")
	wg.Done()
}
