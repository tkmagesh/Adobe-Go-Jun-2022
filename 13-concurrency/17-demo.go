package main

import (
	"fmt"
	"time"
)

func main() {

	primeNoCh := generatePrimes(3, 100)
	for primeNo := range primeNoCh {
		fmt.Println("Prime No : ", primeNo)
	}
	fmt.Println("Done")
}

func generatePrimes(start, end int) <-chan int {
	primeNoCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				time.Sleep(500 * time.Millisecond)
				primeNoCh <- no
			}
			no++
		}
		close(primeNoCh)
	}()
	return primeNoCh
}

func isPrime(n int) bool {
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
