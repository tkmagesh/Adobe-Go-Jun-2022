package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	if count, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println("Invalid data")
		return
	} else {
		for idx := 1; idx <= count; idx++ {
			wg.Add(1)       //increment the wg counter by 1
			go fn(idx, &wg) //scheduling this function execution through the scheduler
		}
		wg.Wait() //block until the wg counter becomes 0
	}
	fmt.Println("Done")
	var input string
	fmt.Scanln(&input)
}

func fn(id int, wg *sync.WaitGroup) {
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(uint64(rand.Intn(5000))) * time.Millisecond)
	fmt.Printf("fn[%d] completed\n", id)
	wg.Done() // decrement the wg counter by 1
}
