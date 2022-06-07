package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("A sample string")
	fmt.Println(str.Length())
}
