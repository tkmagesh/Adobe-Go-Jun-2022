/* converting a panic into an error */
package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("Divisor cannot be zero")

func main() {

	q, r, err := divideClient(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Do not attempt to divide by zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong")
		return
	}
	fmt.Println(q, r)
}

func divideClient(x, y int) (q, r int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = DivideByZeroError
			return
		}
	}()
	q, r = divide(x, y)
	return
}

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
