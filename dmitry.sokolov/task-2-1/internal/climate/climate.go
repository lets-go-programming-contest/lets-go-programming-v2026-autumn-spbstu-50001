package climate

const (
	defaultMaxTemp = 30
	defaultMinTemp = 15
)

type climate struct {
	maxTemp      int
	minTemp      int
	temperatures []int
}

func CalculateClimateInRoom(signs []string, vals []int) []int {
	climateInRoom := climate{
		maxTemp:      defaultMaxTemp,
		minTemp:      defaultMinTemp,
		temperatures: make([]int, len(vals)),
	}

	for index := range len(signs) {
		updateClimate(&climateInRoom, index, vals[index], signs[index])
	}

	return climateInRoom.temperatures
}

func updateClimate(currClimate *climate, index int, preferredTemp int, sign string) {
	switch sign {
	case ">=":
		currClimate.minTemp = max(currClimate.minTemp, preferredTemp)
	case "<=":
		currClimate.maxTemp = min(currClimate.maxTemp, preferredTemp)
	default:

		return
	}

	if currClimate.minTemp > currClimate.maxTemp {
		currClimate.temperatures[index] = -1
	} else {
		currClimate.temperatures[index] = currClimate.minTemp
	}
}
