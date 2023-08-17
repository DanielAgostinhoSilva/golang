package events

import (
	"encoding/json"
	"github.com/google/uuid"
)

type EventSpotProps struct {
	Id        uuid.UUID
	Location  string
	Reserved  bool
	Published bool
}

type EventSpot struct {
	id        EventSpotId
	location  string
	reserved  bool
	published bool
}

func NewEventSpot(props EventSpotProps) (*EventSpot, error) {
	id := *NewEventSpotId(props.Id)
	return &EventSpot{
		id,
		props.Location,
		props.Reserved,
		props.Published,
	}, nil
}

func CreateEventSpot() (*EventSpot, error) {
	eventSpot, err := NewEventSpot(EventSpotProps{
		uuid.New(),
		"",
		false,
		false,
	})
	return eventSpot, err
}

func (props *EventSpot) ChangeLocation(location string) {
	props.location = location
}

func (props *EventSpot) Publish() {
	props.published = true
}

func (props *EventSpot) UnPublish() {
	props.published = false
}

func (props *EventSpot) GetId() uuid.UUID {
	return props.id.value
}

func (props *EventSpot) GetLocation() string {
	return props.location
}

func (props *EventSpot) IsReserved() bool {
	return props.reserved
}

func (props *EventSpot) IsPublished() bool {
	return props.published
}

func (props *EventSpot) ToJson() string {
	customerProps := &EventSpotProps{
		props.id.GetValue(),
		props.location,
		props.reserved,
		props.published,
	}
	jsonProps, _ := json.Marshal(customerProps)
	return string(jsonProps)
}
