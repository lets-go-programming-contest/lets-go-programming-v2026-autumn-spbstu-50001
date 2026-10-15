package main

import (
	"fmt"

	"github.com/dmitrys6/task-2-1/internal/climate"
)

func main() {

	var n, k int

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Incorrect input")
		return
	}

	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Println("Incorrect input")
			return
		}

		t, err := climate.ProcessRoom(k)

		if err != nil {
			fmt.Println("Incorrect input")
			return
		}

		fmt.Println(t)
	}
}
