package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int
LOOP:
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&userChoice)

		switch userChoice {
		case 1:
			fmt.Println("Enter 2 numbers :")
			fmt.Scanln(&n1, &n2)
			result = n1 + n2
		case 2:
			fmt.Println("Enter 2 numbers :")
			fmt.Scanln(&n1, &n2)
			result = n1 - n2
		case 3:
			fmt.Println("Enter 2 numbers :")
			fmt.Scanln(&n1, &n2)
			result = n1 * n2
		case 4:
			fmt.Println("Enter 2 numbers :")
			fmt.Scanln(&n1, &n2)
			result = n1 / n2
		case 5:
			break LOOP
		default:
			fmt.Println("Invalid choice")
			continue LOOP
		}
		fmt.Println("Result =", result)
	}
	fmt.Println("Thank you!")
}
