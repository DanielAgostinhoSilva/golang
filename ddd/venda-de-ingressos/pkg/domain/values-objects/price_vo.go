package values_objects

import "errors"

var (
	ErrInvalidPrice = errors.New("Invalid Price")
)

type Price struct {
	value float64
}

func NewPrice(price float64) (*Price, error) {
	if price < 0.0 {
		return nil, ErrInvalidPrice
	}

	return &Price{value: price}, nil
}

func (props *Price) GetValue() float64 {
	return props.value
}

func (props *Price) Equals(price Price) bool {
	return props.value == price.value
}
