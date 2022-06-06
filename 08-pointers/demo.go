package main

import "fmt"

func main() {
	var no int = 100
	var noPtr *int = &no
	fmt.Println(noPtr)

	//dereferencing - getting the value using the address
	fmt.Println(*noPtr)

	fmt.Println("Before incremening, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	n1, n2 := 10, 20
	fmt.Printf("Before swapping, n1 = %d & n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d & n2 = %d\n", n1, n2)
}

func increment(x *int) {
	*x++
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
