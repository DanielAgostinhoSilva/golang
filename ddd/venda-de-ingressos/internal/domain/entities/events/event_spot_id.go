package events

import "github.com/google/uuid"

type EventSpotId struct {
	value uuid.UUID
}

func NewEventSpotId(id uuid.UUID) *EventSpotId {
	return &EventSpotId{id}
}

func (props EventSpotId) GetValue() uuid.UUID {
	return props.value
}
