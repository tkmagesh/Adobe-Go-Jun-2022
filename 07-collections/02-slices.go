/* Slices */

package main

import "fmt"

func main() {
	//var nos []int
	//var nos []int = []int{}
	//var nos []int = []int{3, 1, 4, 2, 5}
	//var nos = []int{3, 1, 4, 2, 5}
	//nos := []int{3, 1, 4, 2, 5}

	//memory allocated and initilized with the zero-value of int
	/*
		nos := make([]int, 5)
		nos[0] = 3
		nos[1] = 1
		nos[2] = 4
		nos[3] = 2
		nos[4] = 5
		fmt.Println(nos[0])
	*/

	//pre-allocate the memory
	/*
		nos := make([]int, 0, 10)
		fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
		nos = append(nos, 10)
		fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	*/

	//pre-allocate memory and initialize some of them
	nos := make([]int, 2, 10)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos[0] = 10
	nos[1] = 20
	nos = append(nos, 30)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	fmt.Printf("%T\n", nos)
	fmt.Println(nos)

	//append a new value to the slice
	fmt.Println("After appending a new value")
	nos = append(nos, 100)
	fmt.Println(nos)

	//len & cap
	products := []string{"Pen", "Pencil"}
	fmt.Printf("len(products) = %d, cap(products) = %d, products = %v\n", len(products), cap(products), products)

	fmt.Println("After appending a new value")
	products = append(products, "Marker")
	fmt.Printf("len(products) = %d, cap(products) = %d, products = %v\n", len(products), cap(products), products)

	products = append(products, "Scribble-Pad")
	fmt.Printf("len(products) = %d, cap(products) = %d, products = %v\n", len(products), cap(products), products)

	//products = append(products, "Mouse", "Keyboard", "Monitor")
	newProducts := []string{"mouse", "Keyboard", "Monitor"}
	products = append(products, newProducts...)
	fmt.Printf("len(products) = %d, cap(products) = %d, products = %v\n", len(products), cap(products), products)

	fmt.Println("Slicing")
	/*
		slice[lo : hi] => slice[lo] to slice[hi-1]
		slice[lo :] => slice[lo] to end
		slice[:hi] => slice[0] to slice[hi-1]
		slice[:] => copy of the slice
	*/

	fmt.Println("products[2:4] => ", products[2:4])
	fmt.Println("products[:4] => ", products[:4])
	fmt.Println("products[2:] => ", products[2:])

	//creating a copy of the slice
	myProducts := products[:]
	fmt.Println("myProducts", myProducts)
	//modifying the original products
	fmt.Println("After modifying products")
	//products[0] = "mechanical-pencil"
	products = append(products, "sketch-pens")
	fmt.Println("products => ", products)
	fmt.Println("myProducts => ", myProducts)

	fmt.Println()
	fmt.Println("After appending a product to 'myProducts'")
	myProducts = append(myProducts, "ruler")
	fmt.Println("products => ", products)
	fmt.Println("myProducts => ", myProducts)

	fmt.Println()
	//creating a copy of the data
	fmt.Println("Copying the data with the slice")
	dupProducts := make([]string, len(products))
	copy(dupProducts, products)
	fmt.Println("After modifying products")
	products[0] = "artist-pencil"
	fmt.Println("products => ", products)
	fmt.Println("dupProducts => ", dupProducts)

}
