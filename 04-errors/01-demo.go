package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be zero")

func main() {

	q, r, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Dont attempt to divide by zero")
		return
	}
	if err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err)
		return
	}
	fmt.Println(q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, 0, err
	}
	return x / y, x % y, nil
}
*/

/*
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient, remainder = x/y, x%y
	return
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	return
}
