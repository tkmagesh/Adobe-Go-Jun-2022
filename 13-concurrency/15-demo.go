package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10
	primeNoCh := generatePrimes(3, count)
	for count > 0 {
		fmt.Println("Prime No : ", <-primeNoCh)
		count--
	}
	fmt.Println("Done")
}

func generatePrimes(start, count int) <-chan int {
	primeNoCh := make(chan int)
	no := start
	go func() {
		for count > 0 {
			if isPrime(no) {
				time.Sleep(500 * time.Millisecond)
				primeNoCh <- no
				count--
			}
			no++
		}
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
