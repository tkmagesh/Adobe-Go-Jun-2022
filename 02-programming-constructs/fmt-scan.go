package main

import "fmt"

func main() {
	var userName string
	fmt.Println("Enter your name :")
	fmt.Scanln(&userName)
	fmt.Printf("Hi %s, Have a nice day\n", userName)

	var no int
	fmt.Println("Enter a number :")
	fmt.Scanln(&no)
	fmt.Printf("no = %d\n", no)

	var n1, n2 int
	fmt.Println("Enter two numbers :")
	//fmt.Scanln(&n1, &n2)
	fmt.Scanf("%d-%d\n", &n1, &n2)
	fmt.Println("Numbers = ", n1, n2)
}
