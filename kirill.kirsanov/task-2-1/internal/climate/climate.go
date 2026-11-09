package climate

import "errors"

var errEmptyRange = errors.New("temperature range is empty")

type temperatureRange struct {
	low  int
	high int
}

func NewTemperatureRange(minVal, maxVal int) temperatureRange {
	return temperatureRange{
		low:  minVal,
		high: maxVal,
	}
}

func (r *temperatureRange) SetMinLimit(value int) {
	if value > r.low {
		r.low = value
	}
}

func (r *temperatureRange) SetMaxLimit(value int) {
	if value < r.high {
		r.high = value
	}
}

func (r *temperatureRange) GetOptimalTemperature() (int, error) {
	if r.low > r.high {
		return 0, errEmptyRange
	}

	return r.low, nil
}
