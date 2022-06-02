/* higher order functions - functions can be assinged to variables */
package main

import "fmt"

func main() {
	var fn func()
	fn = func() {
		fmt.Println("anon fn(1) invoked")
	}
	fn()

	fn = func() {
		fmt.Println("anon fn(2) invoked")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	fmt.Println(divide(100, 7))
}
