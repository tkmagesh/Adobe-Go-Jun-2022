package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("main started")

	resultCh := add(100, 200)
	result := <-resultCh

	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int) chan int {
	var resultCh chan int = make(chan int)
	go func() {
		fmt.Println("	add operation - started")
		time.Sleep(5 * time.Second)
		result := x + y
		resultCh <- result
		fmt.Println("	add operation - completed")
	}()
	return resultCh
}
