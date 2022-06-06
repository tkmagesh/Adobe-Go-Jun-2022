package main

import (
	"fmt"
	"modularity-demo/calculator"

	//_ "modularity-demo/calculator" //to execute the init function without using any of the apis
	myUtils "modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Modularity demo")

	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("OpCount = ", calculator.OpCount())

	//fmt.Println("Is 21 an even number ? :", utils.IsEven(21))
	fmt.Println("Is 21 an even number ? :", myUtils.IsEven(21))

	color.Red("This line is printed in red color")
	color.Blue("This line is printed in blue color")

}
