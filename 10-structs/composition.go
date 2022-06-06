package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{Product{id, name, cost, units, category}, expiry}
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
	fmt.Println(Foramt(grapes)) //=> Id = 100, Name = "Grapes", Cost = 40, Units = 2, Category = "Food", Expiry = "2 DAYS"

	ApplyDiscount() //APPLY 10% DISCOUNT FOR GRAPES

	fmt.Println(Foramt(grapes)) //=> Id = 100, Name = "Grapes", Cost = 36, Units = 2, Category = "Food", Expiry = "2 DAYS"
}

func Format( /*  */ ) string {

}

func ApplyDiscount( /*  */ ) {

}
