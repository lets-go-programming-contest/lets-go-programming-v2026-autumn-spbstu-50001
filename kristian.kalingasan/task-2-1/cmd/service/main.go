package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	for range N {
		var K int
		min, max := 15, 30
		fmt.Scan(&K)
		for range K {
			var value int
			var op string
			fmt.Scan(&op, &value)
			if op == ">=" && value > min {
				min = value
			} else if op == "<=" && value < max {
				max = value
			}
			if min > max {
				fmt.Println(-1)
			} else {
				fmt.Println(min)
			}
		}
		min = 15
		max = 30
	}
}

/*func main() {
	var x, y int
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&y)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operation string
	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

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
			break
		}
		fmt.Println(x / y)

	default:
		fmt.Println("Invalid operation")
	}
}*/
