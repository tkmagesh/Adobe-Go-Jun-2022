package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

/*
func PrintCircleArea(x Circle) {
	fmt.Println("Area = ", x.Area())
}

func PrintRectangleArea(x Rectangle) {
	fmt.Println("Area = ", x.Area())
}
*/

type ShapeWithArea interface {
	Area() float32
}

func PrintArea(sa ShapeWithArea) {
	fmt.Println("Area = ", sa.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(sp ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", sp.Perimeter())
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintShape(x Shape) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{Radius: 12}
	//fmt.Println("Area = ", c.Area())
	//PrintCircleArea(c)
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/

	PrintShape(c)

	r := Rectangle{Height: 10, Width: 12}
	//fmt.Println("Area = ", r.Area())
	//PrintRectangleArea(r)
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}
