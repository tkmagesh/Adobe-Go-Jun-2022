/* higher order functions - returing functions are return values */

package main

import "fmt"

func main() {
	var fn func() = getFn()
	fn()
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}
