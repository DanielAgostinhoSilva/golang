package entities

import "github.com/google/uuid"

type EventSectionId struct {
	value uuid.UUID
}

func NewEventSectionId(id uuid.UUID) *EventSectionId {
	return &EventSectionId{id}
}

func (props EventSectionId) GetValue() uuid.UUID {
	return props.value
}
