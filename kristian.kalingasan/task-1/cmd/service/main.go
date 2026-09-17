package service

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)
	x, err := strconv.Atoi(input)
	if input[0] == '0' || err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	fmt.Scan(&input)
	y, err := strconv.Atoi(input)
	if input[0] == '0' || err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	var operation string
	fmt.Scan(&operation)
	switch operation {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if y == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(x / y)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
