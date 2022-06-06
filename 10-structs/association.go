package main

import "fmt"

/*
type Employee struct {
	Id   int
	Name string
	City string
	Org  Organization
}
*/

type Employee struct {
	Id   int
	Name string
	City string
	Org  *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	/*
			magesh := Employee{100, "Magesh", "Bangalore", Organization{"Microsoft", "Hydrabad"}}
			fmt.Printf("%#v\n", magesh)

			suresh := Employee{200, "Suresh", "Pune", Organization{"Microsoft", "Hydrabad"}}
			fmt.Printf("%#v\n", suresh)

		fmt.Println("Moving Microsoft from Hydrabad to Delhi")
		magesh.Org.City = "Delhi"
		fmt.Printf("%#v\n", magesh)
		fmt.Printf("%#v\n", suresh)
	*/
	/*
		microsoft := Organization{"Microsoft", "Hydrabad"}

		magesh := Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bangalore",
			Org:  microsoft,
		}
		fmt.Printf("%#v\n", magesh)

		suresh := Employee{
			Id:   200,
			Name: "Suresh",
			City: "Pune",
			Org:  microsoft,
		}
		fmt.Printf("%#v\n", suresh)

		fmt.Println("Moving Microsoft from Hydrabad to Delhi")
		microsoft.City = "Delhi"

		fmt.Printf("%#v\n", magesh)
		fmt.Printf("%#v\n", suresh)
	*/

	//microsoft := &Organization{"Microsoft", "Hydrabad"}
	microsoft := new(Organization) // === microsoft := &Organization{}
	microsoft.Name = "Microsoft"
	microsoft.City = "Hyderabad"

	magesh := Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bangalore",
		Org:  microsoft,
	}
	fmt.Printf("%#v\n", magesh)

	suresh := Employee{
		Id:   200,
		Name: "Suresh",
		City: "Pune",
		Org:  microsoft,
	}
	fmt.Printf("%#v\n", suresh)

	fmt.Println("Moving Microsoft from Hydrabad to Delhi")
	microsoft.City = "Delhi"

	fmt.Printf("%#v\n", magesh.Org.City)
	fmt.Printf("%#v\n", suresh.Org.City)

}
