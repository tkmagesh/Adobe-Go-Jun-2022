/* higher order functions - functions passed as arguments to other functions */
package main

import "fmt"

func main() {
	exec(fn)

	f1 := func() {
		fmt.Println("f1 invoked")
	}
	exec(f1)

	exec(func() {
		fmt.Println("anon function invoked")
	})
}

func fn() {
	fmt.Println("fn invoked")
}

func exec(f func()) {
	f()
}
