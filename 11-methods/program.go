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

func (p *Product) ApplyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

//
func ApplyDiscountFunc(p *Product, discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

func main() {
	pen := Product{105, "Pen", 5, 50, "Stationary"}
	fmt.Println(pen.Name)
	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())
	fmt.Println("After applying 10% discount")
	(&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10)
	//ApplyDiscountFunc(pen, 10)

	fmt.Println(pen.Format())
	pen.WhoAmI()

	pencil := &Product{107, "Pencil", 5, 50, "Stationary"}
	fmt.Println(pencil.Name)
	fmt.Println(pencil.Format())

}
