package main

import "fmt"

func main() {
	/* language constructs for variable declaration => var, :=  */

	/*
		var msg string
		msg = "Hello World!"
	*/
	/*
		var msg string = "Hello World!"
	*/

	/*
		var msg = "Hello World!" //=> type inference
	*/

	msg := "Hello World" // := syntax is applicable ONLY in the function scope (NOT in package scope)
	fmt.Println(msg)

	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Add result ="
		result = x + y
		fmt.Println(x, y, str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Add result ="
		var result int = x + y
		fmt.Println(x, y, str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Add result ="
		result = x + y
		fmt.Println(x, y, str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Add result ="
		result = x + y
		fmt.Println(x, y, str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Add result ="
			result int    = x + y
		)
		fmt.Println(x, y, str, result)
	*/

	/*
		var x, y = 100, 200
		var str = "Add result ="
		var result = x + y

		fmt.Println(x, y, str, result)
	*/

	/*
		var x, y, str = 100, 200, "Add result ="
		var result = x + y
		fmt.Println(x, y, str, result)
	*/

	x, y, str := 100, 200, "Add result ="
	result := x + y

	fmt.Println(x, y, str, result)

	var zeroint int
	var zerobool bool
	var zerostr string
	var zerobyte byte
	fmt.Println(zeroint, zerobool, zerostr, zerobyte)
	fmt.Printf("zerostr = %q\n", zerostr)

	fmt.Println("Constants")
	const pi float32 = 3.14
	fmt.Println("pi =", pi)

	fmt.Println("iota")
	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/
	/*
		const (
			red = iota
			green
			blue
		)
	*/
	/*
		const (
			red = iota + 2
			green
			blue
		)
	*/
	/*
		const (
			red = iota * 2
			green
			blue
		)
	*/
	const (
		red = iota * 2
		green
		_
		blue
	)
	fmt.Println("red =", red, "green =", green, "blue =", blue)

	const (
		read   = 1 << iota // 00000001 = 1
		write              // 00000010 = 2
		remove             // 00000100 = 4

		// admin will have all of the permissions
		admin = read | write | remove
	)
	fmt.Printf("read = %v\n", read)
	fmt.Printf("write = %v\n", write)
	fmt.Printf("remove = %v\n", remove)
	fmt.Printf("admin = %v\n", admin)
}
