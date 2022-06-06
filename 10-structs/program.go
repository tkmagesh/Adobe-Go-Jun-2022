package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{}
	//var emp = Employee{}
	//emp := Employee{}
	//emp := Employee{100, "Magesh", "Bangalore"}
	//emp := Employee{Id: 100, Name: "Magesh"}
	emp := Employee{
		Id:   100,
		Name: "Magesh",
	}
	fmt.Printf("%#v\n", emp)

	//accessing the fields
	fmt.Println("Emp Name = ", emp.Name)

	//using pointers
	newEmp := &Employee{
		Id:   200,
		Name: "Suresh",
	}
	fmt.Println(newEmp)
	//fmt.Println((*newEmp).Name)
	fmt.Println(newEmp.Name)
}
