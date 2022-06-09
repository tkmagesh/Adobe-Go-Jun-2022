package main

import (
	"fmt"
	"time"
)

func main() {
	var resultCh chan int = make(chan int)
	fmt.Println("main started")

	go add(100, 200, resultCh)
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

}
