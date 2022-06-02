package main

import "fmt"

func main() {
	fn()
	greet("Magesh")
	fmt.Println(getGreetMsg("Suresh"))
	fmt.Println("Adding 100 and 200, result =", add(100, 200))
	fmt.Println("Subtract 200 from 100, result =", subtract(100, 200))
	fmt.Println("Multiplying 100 and 200, result =", multiply(100, 200))
	//fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/
	quotient, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d \n", quotient)
}

func fn() {
	fmt.Println("fn invoked")
}

/* function with argument */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* function with argument and return value */
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

/* function with more than one arguments */
func add(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

/* function with named return value */
func multiply(x, y int) (result int) {
	result = x * y
	return
}

/* function returning more than one result */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
