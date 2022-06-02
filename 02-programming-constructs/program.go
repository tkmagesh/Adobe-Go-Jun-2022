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

	fmt.Println("switch case")
	/*
		score 0-3 => "Terrible"
		score 4-7 => "Not Bad"
		score 8-9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/
	score := 7
	/*
		switch score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")
		}
	*/

	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	no := 12
	switch {
	case no%3 == 0:
		fmt.Printf("%d is divisible by 3\n", no)
	case no%4 == 0:
		fmt.Printf("%d is divisible by 4\n", no)
	}

	fmt.Println("switch (with fallthrough)")
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
		fallthrough
	case 7:
		fmt.Println("n <= 7")
		//fallthrough
	case 8:
		fmt.Println("n <= 8")
		fallthrough
	case 9:
		fmt.Println("n <= 9")
	}

	subscription := "pro"
	switch subscription {
	case "premium":
		fmt.Println("All <premium> features are included")
		fallthrough
	case "pro":
		fmt.Println("All <pro> features are included")
		fallthrough
	case "basic":
		fmt.Println("All <basic> features are included")
		fallthrough
	case "free":
		fmt.Println("All <free> features are included")

	}
}
