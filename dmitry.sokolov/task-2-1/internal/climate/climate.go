package climate

import "fmt"

func ProcessRoom(employees int) (int, error) {
	var (
		sign                   string
		curTemp, preferredTemp int
		maxTemp                = 30
		minTemp                = 15
	)

	for preference := range employees {
		_, err := fmt.Scan(&sign)

		if err != nil {
			fmt.Println("Incorrect input")

			return curTemp, fmt.Errorf("Failed to scan sign: %w", err)
		}

		_, err = fmt.Scan(&preferredTemp)

		if err != nil {
			fmt.Println("Incorrect input")

			return curTemp, fmt.Errorf("Failed to scan temperature: %w", err)
		}

		if preference == 0 {
			curTemp = preferredTemp
		}

		curTemp, maxTemp, minTemp = updateTemp(
			maxTemp, minTemp, curTemp, preferredTemp, sign)
	}

	return curTemp, nil
}

func updateTemp(maxTemp int, minTemp int, curTemp int,
	preferredTemp int, sign string) (int, int, int) {

	switch sign {
	case ">=":
		minTemp = max(minTemp, preferredTemp)
	case "<=":
		maxTemp = min(maxTemp, preferredTemp)
	}

	if minTemp > maxTemp {
		return -1, maxTemp, minTemp
	}

	if curTemp < minTemp || curTemp > maxTemp {
		curTemp = preferredTemp
	}

	return curTemp, maxTemp, minTemp
}
