package main

import "fmt"

func main() {
	//var x interface{}
	var x any
	//x = 100
	//x = "This is a string"
	//x = true
	//x = 10.786
	//x = struct{}{}
	x = []int{3, 1, 4, 2, 5}
	fmt.Println("x =", x)
	//fmt.Println("2x =", 2*x)

	//type assertion
	if val, ok := x.(int); ok {
		fmt.Println("2x =", 2*val)
	} else {
		fmt.Println("x is not an int")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("2x = ", 2*val)
	case string:
		fmt.Println("length of x = ", len(val))
	case bool:
		fmt.Println("!x = ", !val)
	case struct{}:
		fmt.Println("x is an empty object")
	case []int:
		fmt.Println("x is a list of ints", val)
	default:
		fmt.Println("unknown type!")
	}
}
