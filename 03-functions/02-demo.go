/* anonymous functions */
package main

import "fmt"

func main() {
	func() {
		fmt.Println("anon fn invoked")
	}()

	func(x, y int) {
		fmt.Println("Add Result (anon fn) = ", x+y)
	}(100, 200)

	/*
		subtrctResult := func(x, y int) int {
			return x - y
		}(100, 200)
	*/
	subtrctResult := func(x, y int) (result int) {
		result = x - y
		return
	}(100, 200)
	fmt.Println("Subtracting 200 from 100 (anon fn) = ", subtrctResult)

	quotient, remainder := func(x, y int) (q, r int) {
		/*
			q = x / y
			r = x % y
		*/
		q, r = x/y, x%y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7 (anon fn), quotient = %d and remaider = %d\n", quotient, remainder)
}
