package main

import "fmt"

func main() {
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos [5]int = [5]int{3, 1, 4}
	//nos := [5]int{}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//iterating (using indexer)
	fmt.Println("iterating (using indexer)")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	//iterating (using range)
	fmt.Println("iterating (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	//iterating (using range) (only values)
	fmt.Println("iterating (using range)")
	for _, val := range nos {
		fmt.Println(val)
	}

	newNos := nos
	newNos[0] = 100
	fmt.Println(nos, newNos)

}
