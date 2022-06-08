package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling this function execution through the scheduler
	f2()
	time.Sleep(5 * time.Millisecond)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Millisecond)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
