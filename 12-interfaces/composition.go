package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}
*/

//implementation of the fmt.Stringer interface
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{Product{id, name, cost, units, category}, expiry}
}

//overriding the Product.Format() method

/*
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%v, Expiry = %q", pp.Product.Format(), pp.Expiry)
}
*/

//implementation of the fmt.Stringer interface
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%v, Expiry = %q", pp.Product, pp.Expiry)
}

func main() {
	/* var grapes = PerishableProduct{
		Product: Product{
			Id:       200,
			Name:     "Grapes",
			Cost:     40,
			Units:    2,
			Category: "Food",
		},
		Expiry: "2 DAYS",
	} */

	var grapes = NewPerishableProduct(200, "Grapes", 40, 2, "Food", "2 DAYS")

	//fmt.Println(grapes.Product.Name)
	//fmt.Println(Format(*grapes)) //=> Id = 100, Name = "Grapes", Cost = 40, Units = 2, Category = "Food", Expiry = "2 DAYS"
	//fmt.Println(grapes.Format())
	fmt.Println(grapes)

	//ApplyDiscount(grapes, 10) //APPLY 10% DISCOUNT FOR GRAPES
	grapes.ApplyDiscount(10)

	//fmt.Println(Format(*grapes)) //=> Id = 100, Name = "Grapes", Cost = 36, Units = 2, Category = "Food", Expiry = "2 DAYS"
	//fmt.Println(grapes.Format())
	fmt.Println(grapes)
}

/*  */
