package climate

import "errors"

var ErrEmptyRange = errors.New("temperature range is empty")

type TemperatureRange struct {
	low  int
	high int
}

func NewTemperatureRange(min, max int) TemperatureRange {
	return TemperatureRange{
		low:  min,
		high: max,
	}
}

func (r *TemperatureRange) SetMinLimit(value int) {
	if value > r.low {
		r.low = value
	}
}

func (r *TemperatureRange) SetMaxLimit(value int) {
	if value < r.high {
		r.high = value
	}
}

func (r *TemperatureRange) GetOptimalTemperature() (int, error) {
	if r.low > r.high {
		return 0, ErrEmptyRange
	}
	return r.low, nil
}
