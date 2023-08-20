package entities

import (
	"github.com/google/uuid"
)

type CustomerId struct {
	value uuid.UUID
}

func NewCustomerId(id uuid.UUID) *CustomerId {
	return &CustomerId{id}
}

func (props *CustomerId) GetValue() uuid.UUID {
	return props.value
}
