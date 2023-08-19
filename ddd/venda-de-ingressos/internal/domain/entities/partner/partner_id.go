package events

import "github.com/google/uuid"

type PartnerId struct {
	value uuid.UUID
}

func NewPartnerId(id uuid.UUID) *PartnerId {
	return &PartnerId{id}
}

func (props PartnerId) GetValue() uuid.UUID {
	return props.value
}
