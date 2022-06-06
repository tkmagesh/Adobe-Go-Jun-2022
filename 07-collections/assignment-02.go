package main

import "fmt"

var operations map[int]func(int, int) int = map[int]func(int, int) int{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	var userChoice, n1, n2 int
	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}
		if _, exists := operations[userChoice]; !exists {
			fmt.Println("Invalid choice")
			continue
		}
		n1, n2 = getOperands()
		performOperation(userChoice, n1, n2)
	}
	fmt.Println("Thank you!")
}

func performOperation(userChoice, n1, n2 int) {
	result := operations[userChoice](n1, n2)
	fmt.Println("Result =", result)
}
func getOperands() (int, int) {
	var n1, n2 int
	fmt.Println("Enter 2 numbers :")
	fmt.Scanln(&n1, &n2)
	return n1, n2
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice:")
	fmt.Scanln(&userChoice)
	return userChoice
}
func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
