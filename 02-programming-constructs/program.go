package main

import "fmt"

func main() {
	/* if */

	/*
		no := 21
		if no%2 == 0 {
			fmt.Println("no is an even number")
		} else {
			fmt.Println("no is an odd number")
		}
		fmt.Println("no = ", no)
	*/

	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	/* for (v1.0) */
	fmt.Println("for (v1.0)")
	for idx := 1; idx <= 10; idx++ {
		fmt.Println(idx)
	}

	/* for (while) (v2.0) */
	fmt.Println("for (while) (v2.0)")
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %d\n", numSum)

	/* for (infinite) (v3.0) */
	fmt.Println("for (infinite) (v3.0)")
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println("x = ", x)

	/* using labels with loops */
	fmt.Println("using labels with loops")
LOOP:
	for i := 1; i <= 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d%d\n", i, j)
			if i == j {
				fmt.Println("------")
				continue LOOP
			}
		}
	}
}
