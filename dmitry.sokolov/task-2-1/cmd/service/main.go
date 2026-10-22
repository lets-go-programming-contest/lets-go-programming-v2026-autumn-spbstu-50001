package main

import (
	"fmt"

	"github.com/dmitrys6/task-2-1/internal/climate"
)

func main() {
	var numberOfRooms, employeesInRoom int

	_, err := fmt.Scan(&numberOfRooms)
	if err != nil || numberOfRooms < 0 {
		fmt.Println("Incorrect input")

		return
	}

	for range numberOfRooms {
		_, err = fmt.Scan(&employeesInRoom)
		if err != nil || employeesInRoom < 0 {
			fmt.Println("Incorrect input")

			return
		}

		var (
			sign          string
			preferredTemp int
			signs         = make([]string, employeesInRoom)
			temps         = make([]int, employeesInRoom)
		)

		for index := range employeesInRoom {
			_, err := fmt.Scan(&sign)
			if err != nil || (sign != ">=" && sign != "<=") {
				fmt.Println("Incorrect input")

				return
			}

			signs[index] = sign

			_, err = fmt.Scan(&preferredTemp)
			if err != nil {
				fmt.Println("Incorrect input")

				return
			}

			temps[index] = preferredTemp
		}
		minTemps := climate.CalculateClimateInRoom(signs, temps)

		for _, val := range minTemps {
			fmt.Println(val)
		}
	}
}
