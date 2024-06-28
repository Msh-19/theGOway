package main

import (
	"fmt"
	"os"
	// "strconv"
)

func main(){

	var num1, num2 float64
	var operator string

	fmt.Print("Enter first Number: ")
	fmt.Scanln(&num1)
	
	fmt.Print("Enter the operator: ")
	fmt.Scanln(&operator)

	fmt.Print("Enter Second Number: ")
	fmt.Scanln(&num2)

	result,err := calculate(num1,num2,operator)
	if err != nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f\n", result)

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