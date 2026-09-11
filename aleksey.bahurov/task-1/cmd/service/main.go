package main

import "fmt"

func main() {
	var (
		firstOperand, secondOperand int
		operation                   string
	)

	_, error := fmt.Scanln(&firstOperand)
	if error != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, error = fmt.Scanln(&secondOperand)
	if error != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, error = fmt.Scanln(&operation)
	if error != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(firstOperand + secondOperand)
	case "-":
		fmt.Println(firstOperand - secondOperand)
	case "*":
		fmt.Println(firstOperand * secondOperand)
	case "/":
		if secondOperand == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(firstOperand / secondOperand)
	default:
		fmt.Println("Invalid operation")
	}
}
