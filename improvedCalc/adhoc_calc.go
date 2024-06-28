package main

import(
	"fmt"
	"os"
)

func main(){
	for{
	var num1 float64
	var num2 float64

	var operator string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil{
		fmt.Println("Error reading first number:",err)
		os.Exit(1)
	}

	fmt.Print("Enter the operator (+,-,*,/): ")
	_, err = fmt.Scanln(&operator)
	if err != nil{
		fmt.Println("Error reading operator:",err)
		os.Exit(1)
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Error reading second number:",err)
		os.Exit(1)
	}

	// where calculation happens
	result, err := calculate(num1, num2, operator)
	if err != nil{
		fmt.Println("Error :",err)
	}else {
		fmt.Printf("Result: %.2f\n", result)
	}

	fmt.Print("Do you want to continue? (yes/press anything): ")

	var choice string
	_,err = fmt.Scanln(&choice)
	if err != nil{
		fmt.Println("Error reading choice:",err)
		os.Exit(1)
	}
	if choice != "yes"{
		break
	}
}
	fmt.Println("Goodbye! thanks for using our calculator")
}

func calculate(num1, num2 float64, operator string)(float64, error){
	switch operator{
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}