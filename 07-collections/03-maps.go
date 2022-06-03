/* maps */

package main

import "fmt"

func main() {
	//var productsRank map[string]int
	/*
		var productsRank map[string]int = make(map[string]int)
		productsRank["Pen"] = 1
	*/
	/*
		productsRank := map[string]int{}
		productsRank["Pen"] = 1
	*/

	//productsRank := map[string]int{"Pen": 2, "Pencil": 1, "Marker": 4}
	productsRank := map[string]int{
		"Pen":    2,
		"Pencil": 1,
		"Marker": 4,
	}
	fmt.Println(productsRank)

	fmt.Println("Adding a new item")
	productsRank["stappler"] = 5
	fmt.Println(productsRank)

	fmt.Println()
	fmt.Println("Iterating a map")
	for key, val := range productsRank {
		fmt.Printf("Key = %s, val = %d\n", key, val)
	}

	fmt.Println("Checking if a key exists")
	keyToCheck := "Mouse"
	if rank, exists := productsRank[keyToCheck]; exists {
		fmt.Printf("Rank of %s is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key [%s] doesnot exist\n", keyToCheck)
	}

	fmt.Println("Removing a key")
	delete(productsRank, "Mouse")
	fmt.Println(productsRank)
}
