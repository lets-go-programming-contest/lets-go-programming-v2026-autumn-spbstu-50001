package climate

import "errors"

var (
	ErrInvalidOperator = errors.New("invalid comparison operator")
	ErrEmptyRange      = errors.New("temperature range is empty")
)

type Constraint struct {
	isLowerBound bool
	value        int
}

func NewConstraint(cmp string, temp int) (Constraint, error) {
	if cmp != ">=" && cmp != "<=" {
		return Constraint{}, ErrInvalidOperator
	}
	return Constraint{isLowerBound: cmp == ">=", value: temp}, nil
}

type TemperatureRange struct {
	low  int
	high int
}

func NewTemperatureRange(low, high int) TemperatureRange {
	return TemperatureRange{low: low, high: high}
}

func (r *TemperatureRange) Apply(c Constraint) {
	if c.isLowerBound {
		if c.value > r.low {
			r.low = c.value
		}
		return
	}
	if c.value < r.high {
		r.high = c.value
	}
}

func (r *TemperatureRange) GetOptimal() (int, error) {
	if r.low > r.high {
		return -1, ErrEmptyRange
	}
	return r.low, nil
}
