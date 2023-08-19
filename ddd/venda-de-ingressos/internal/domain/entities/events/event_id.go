package events

import "github.com/google/uuid"

type EventId struct {
	value uuid.UUID
}

func NewEventId(id uuid.UUID) *EventId {
	return &EventId{id}
}

func (props EventId) GetValue() uuid.UUID {
	return props.value
}
