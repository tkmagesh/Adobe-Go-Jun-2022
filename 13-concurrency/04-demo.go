package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)  //increment the wg counter by 1
	go f1(&wg) //scheduling this function execution through the scheduler
	f2()
	wg.Wait() //block until the wg counter becomes 0
	fmt.Println("Done")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
