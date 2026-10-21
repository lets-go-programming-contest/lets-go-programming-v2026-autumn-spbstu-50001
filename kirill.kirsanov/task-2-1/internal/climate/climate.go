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

func (r *TemperatureRange) Apply(constraint Constraint) {
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

func (r *TemperatureRange) GetOptimal() (int, error) {
	if r.low > r.high {
		return -1, ErrEmptyRange
	}

	return r.low, nil
}
