package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	for i := 0; i < 2; i++ {
		select {
		case ch1No := <-ch1:
			fmt.Println("data from ch1:", ch1No)
		case ch2No := <-ch2:
			fmt.Println("data from ch2:", ch2No)
			/* default:
			fmt.Println("No data from either channels") */
		}
	}

}
