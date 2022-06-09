package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int)

	go func() {
		fmt.Println("goroutine started")
		time.Sleep(5 * time.Second)
		ch <- 100
		fmt.Println("goroutine completed")
	}()

	data := <-ch
	fmt.Println("data from channel = ", data)
}
