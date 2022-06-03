package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("application exits abruptly due to a panic")
			return
		}
		fmt.Println("successful execution of the app!")
	}()
	q, r := divide(100, 7)
	fmt.Println(q, r)
}

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
