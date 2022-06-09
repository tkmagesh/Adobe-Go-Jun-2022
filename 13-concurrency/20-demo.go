package main

import (
	"fmt"
	"time"
)

func main() {

	primeNoCh := generatePrimes(3)
	for primeNo := range primeNoCh {
		fmt.Println("Prime No : ", primeNo)
	}
	fmt.Println("Done")
}

func generatePrimes(start int) <-chan int {
	primeNoCh := make(chan int)
	timeOutCh := time.After(20 * time.Second)
	go func() {
		no := start
	LOOP:
		for {
			if !isPrime(no) {
				no++
				select {
				case <-timeOutCh:
					break LOOP
				default:
					continue LOOP
				}
			}
			select {
			case primeNoCh <- no:
				time.Sleep(500 * time.Millisecond)
				no++
			case <-timeOutCh:
				break LOOP
			}
		}
		close(primeNoCh)
	}()
	return primeNoCh
}

func timeOut(d time.Duration) <-chan time.Time {
	timeOutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeOutCh <- time.Now()
	}()
	return timeOutCh
}

func isPrime(n int) bool {
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
