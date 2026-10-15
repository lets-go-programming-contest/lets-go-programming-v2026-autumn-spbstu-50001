package climate

import "fmt"

func ProcessRoom(employees int) error {
	var (
		sign                   string
		curTemp, preferredTemp int
		maxTemp                = 30
		minTemp                = 15
	)

	for range employees {
		_, err := fmt.Scan(&sign)
		if err != nil {
			return fmt.Errorf("failed to scan sign: %w", err)
		}

		_, err = fmt.Scan(&preferredTemp)
		if err != nil {
			return fmt.Errorf("failed to scan temperature: %w", err)
		}

		curTemp, maxTemp, minTemp = updateTemp(
			maxTemp, minTemp, preferredTemp, sign)

		fmt.Println(curTemp)
	}

	return nil
}

func updateTemp(maxTemp int, minTemp int, preferredTemp int, sign string) (int, int, int) {
	switch sign {
	case ">=":
		minTemp = max(minTemp, preferredTemp)
	case "<=":
		maxTemp = min(maxTemp, preferredTemp)
	}

	if minTemp > maxTemp {
		return -1, maxTemp, minTemp
	}

	return minTemp, maxTemp, minTemp
}
