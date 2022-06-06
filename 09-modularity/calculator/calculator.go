package calculator

import "fmt"

var opCount int

func init() {
	fmt.Println("Calculator initialized")
}

func OpCount() int {
	return opCount
}
