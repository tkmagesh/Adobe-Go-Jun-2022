package main

import (
	"fmt"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func Format(p Product) string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/* write the functions for the following :

        IndexOf => return the index of the given product
            ex: IndexOf(products, product) => returns the index of the given product

        Includes => return true if the given product is present in the collection else return false
            ex: Includes(products, product) => returns true if the given product is present in the collection

        Filter => return a new collection of products that satisfy the given condition
            use cases:
                1. filter all costly products (cost > 1000)
					OR
                2. filter all stationary products (category = "Stationary")
					OR
                3. filter all costly stationary products
				etc

        Any => return true if any of the product in the collections satifies the given criteria
            use cases:
                1. are there any costly product (cost > 1000)?
                2. are there any stationary product (category = "Stationary")?
                3. are there any costly stationary product?
				etc

        All => return true if all the products in the collections satifies the given criteria
            use cases:
                1. are all the products costly products (cost > 1000)?
                2. are all the products stationary products (category = "Stationary")?
                3. are all the products costly stationary products?
				etc */

func indexOf(products []Product, product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func includes(products []Product, product Product) bool {
	for _, p := range products {
		if p == product {
			return true
		}
	}
	return false
}

/* func filterCostlyProducts(products []Product) []Product {
	result := []Product{}
	for _, p := range products {
		if p.Cost > 1000 {
			result = append(result, p)
		}
	}
	return result
}

func filterStationaryProducts(products []Product) []Product {
	result := []Product{}
	for _, p := range products {
		if p.Category == "Stationary" {
			result = append(result, p)
		}
	}
	return result
} */

func filter(products []Product, predicate func(Product) bool) []Product {
	result := []Product{}
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func any(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

func all(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

func FormatProducts(products []Product) string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%s\n", Format(p))
	}
	return result
}

func main() {
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	/*
		p1 := Product{105, "Pen", 5, 50, "Stationary"}
		p2 := Product{105, "Pen", 5, 50, "Stationary"}
		fmt.Println(p1 == p2)
	*/

	costlyProductPredicate := func(p Product) bool {
		return p.Cost > 1000
	}
	costlyProducts := filter(products, costlyProductPredicate)
	fmt.Println("Costly Products")
	fmt.Println(FormatProducts(costlyProducts))

	fmt.Println("Stationary Products")
	stationaryProducts := filter(products, func(p Product) bool {
		return p.Category == "Stationary"
	})
	fmt.Println(FormatProducts(stationaryProducts))

	fmt.Println("Are there any costly products ? : ", any(products, costlyProductPredicate))

	fmt.Println("Are all the products costly products ? : ", all(products, costlyProductPredicate))
}
