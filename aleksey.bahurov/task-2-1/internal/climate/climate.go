package climate

import "errors"

var (
	ErrInvalidOperator = errors.New("invalid comparison operator")
	ErrInvalidRange    = errors.New("invalid temperature range")
	ErrEmptyRange      = errors.New("temperature range is empty")
)

type constraint struct {
	isLowerBound bool
	value        int
}

func NewConstraint(cmp string, temp int) (constraint, error) {
	switch cmp {
	case ">=":
		return constraint{isLowerBound: true, value: temp}, nil
	case "<=":
		return constraint{isLowerBound: false, value: temp}, nil
	default:
		return constraint{}, ErrInvalidOperator
	}
}

type temperatureRange struct {
	low  int
	high int
}

func NewTemperatureRange(low, high int) (temperatureRange, error) {
	if low > high {
		return temperatureRange{}, ErrInvalidRange
	}

	return temperatureRange{low: low, high: high}, nil
}

func (r *temperatureRange) ApplyConstraint(constraint constraint) {
	if constraint.isLowerBound {
		if constraint.value > r.low {
			r.low = constraint.value
		}

		return
	}

	if constraint.value < r.high {
		r.high = constraint.value
	}
}

func (r *temperatureRange) GetOptimalTemperature() (int, error) {
	if r.low > r.high {
		return 0, ErrEmptyRange
	}

	return r.low, nil
}
