package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p Product) WhoAmI() {
	fmt.Printf("I am a Product [%q]\n", p.Name)
	//fmt.Println("I am a Product")
}

func main() {
	pen := Product{105, "Pen", 5, 50, "Stationary"}

	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())

	pen.WhoAmI()

}
