package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("Divisor cannot be zero")

func main() {
	defer func() {
		e := recover()
		if e != nil && e.(error) == DivideByZeroError {
			fmt.Println("Do not attempt to divide by zero")
			return
		}
		if e != nil {
			fmt.Println("application exits abruptly due to a panic")
			return
		}
		fmt.Println("successful execution of the app!")
	}()
	q, r := divide(100, 0)
	fmt.Println(q, r)
}

func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient, remainder = x/y, x%y
	return
}
