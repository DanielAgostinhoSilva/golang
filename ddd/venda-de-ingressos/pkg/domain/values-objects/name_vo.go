package values_objects

import "errors"

var (
	ErrNameIsRequired = errors.New("name is required")
	ErrInvalidName    = errors.New("Invalid NAME")
)

type Name struct {
	value string
}

func NewName(name string) (*Name, error) {
	if name == "" {
		return nil, ErrNameIsRequired
	}
	if len(name) < 3 {
		return nil, ErrInvalidName
	}
	return &Name{value: name}, nil
}

func (props *Name) GetValue() string {
	return props.value
}

func (props *Name) Equals(name Name) bool {
	return props.value == name.value
}
