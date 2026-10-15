package climate

import "fmt"

func ProcessRoom(k int) (int, error) {

	var (
		sign                   string
		curTemp, preferredTemp int
		maxTemp                int = 30
		minTemp                int = 15
	)

	for j := range k {
		_, err := fmt.Scan(&sign)
		if err != nil {
			fmt.Println("Incorrect input")
			return curTemp, err
		}

		_, err = fmt.Scan(&preferredTemp)
		if err != nil {
			fmt.Println("Incorrect input")
			return curTemp, err
		}

		if j == 0 {
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
