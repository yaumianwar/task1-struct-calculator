package main

import (
	"fmt"

	calc "github.com/yaumianwar/calculator/Calc"
)

func main() {
	const (
		ADD       = "+"
		SUBSTRACT = "-"
		MULTIPLY  = "*"
		DIVIDE    = "/"
	)

	calculator := calc.Calc{}

	// input the first number
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%f", &calculator.Data1)

	// input the second number
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%f", &calculator.Data2)

	// input operator
	fmt.Print("Enter operator (+,-,*,/): ")
	fmt.Scanf("%s", &calculator.Opr)

	switch calculator.Opr {
	case ADD:
		fmt.Println(calculator.AddNumber())
	case SUBSTRACT:
		fmt.Println(calculator.SubstractNumber())
	case MULTIPLY:
		fmt.Println(calculator.MultiplyNumber())
	case DIVIDE:
		fmt.Println(calculator.DivideNumber())
	default:
		fmt.Println("Please Enter a correct operator")
	}

}
