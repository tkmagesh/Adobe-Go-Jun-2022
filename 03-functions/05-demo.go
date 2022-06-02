/* higher order functions - functions passed as arguments to other functions */
package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
	logOperation(100, 200, func(x, y int) {
		fmt.Println("Multiply result =", x*y)
	})
}

func logOperation(x, y int, oper func(int, int)) {
	fmt.Println("Before invocation")
	oper(x, y)
	fmt.Println("After invocation")
}

func logAdd(x, y int) {
	fmt.Println("Before invocation")
	add(x, y)
	fmt.Println("After invocation")
}

func logSubtract(x, y int) {
	fmt.Println("Before invocation")
	subtract(x, y)
	fmt.Println("After invocation")
}

/* assumption : the following function code cannot be changed */
func add(x, y int) {
	fmt.Println("Add result =", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result =", x-y)
}
