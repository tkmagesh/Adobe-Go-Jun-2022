package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
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
				etc

		Sort => Sort the products collection by any attribute
			IMPORTANT : Use sort.Sort() function
            use cases:
                1. sort the products collection by cost
                2. sort the products collection by name
                3. sort the products collection by units
                4. sort the products collection by cost in descending order
                5. sort the products collection by name in descending order
                6. sort the products collection by units in descending order
				etc

*/

type Products []Product

func (products Products) IndexOf(product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	for _, p := range products {
		if p == product {
			return true
		}
	}
	return false
}

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

func (products Products) String() string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p)
	}
	return result
}

//sort.Interface implementation
func (products Products) Len() int {
	return len(products)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	}
}

func main() {
	products := Products{
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
	//costlyProducts := filter(products, costlyProductPredicate)
	costlyProducts := products.Filter(costlyProductPredicate)

	fmt.Println("Costly Products")
	//fmt.Println(FormatProducts(costlyProducts))
	fmt.Println(costlyProducts)

	fmt.Println("Stationary Products")
	stationaryProducts := products.Filter(func(p Product) bool {
		return p.Category == "Stationary"
	})
	//fmt.Println(FormatProducts(stationaryProducts))
	fmt.Println(stationaryProducts)

	//fmt.Println("Are there any costly products ? : ", any(products, costlyProductPredicate))
	fmt.Println("Are there any costly products ? : ", products.Any(costlyProductPredicate))

	//fmt.Println("Are all the products costly products ? : ", all(products, costlyProductPredicate))
	fmt.Println("Are all the products costly products ? : ", products.All(costlyProductPredicate))

	fmt.Println()
	fmt.Println("Sorting")
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Default sort (by id)")
	sort.Sort(products)
	fmt.Println(products)

	fmt.Println("Sort by Name")
	//sort.Sort(ByName{products})
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println("Sort by Cost")
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println("Sort by Units")
	products.Sort("Units")
	fmt.Println(products)

	fmt.Println("Sort by Category")
	products.Sort("Category")
	fmt.Println(products)
}
