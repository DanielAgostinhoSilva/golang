package entities

import (
	"github.com/google/uuid"
	"time"
	values_objects "venda-de-ingressos/pkg/domain/values-objects"
)

type CreateEventCommand struct {
	Name        string
	Description string
	Date        time.Time
	PartnerId   uuid.UUID
}

type AddSectionCommand struct {
	Name        string
	Description string
	TotalSpot   int
	Price       float64
}

type EventProps struct {
	Id                uuid.UUID
	Name              string
	Description       string
	Date              time.Time
	Published         bool
	TotalSpots        int
	TotalSpotReserved int
	PartnerId         uuid.UUID
	Sections          []EventSection
}

type Event struct {
	id                EventId
	name              values_objects.Name
	description       string
	date              time.Time
	published         bool
	totalSpots        int
	totalSpotReserved int
	partnerId         PartnerId
	sections          []EventSection
}

func NewEvent(props EventProps) (*Event, error) {
	name, err := values_objects.NewName(props.Name)
	return &Event{
		id:                *NewEventId(props.Id),
		name:              *name,
		description:       props.Description,
		date:              props.Date,
		published:         props.Published,
		totalSpots:        props.TotalSpots,
		totalSpotReserved: props.TotalSpotReserved,
		partnerId:         *NewPartnerId(props.PartnerId),
		sections:          props.Sections,
	}, err
}

func CreateEvent(command CreateEventCommand) (*Event, error) {
	props := &EventProps{
		Id:          uuid.New(),
		Name:        command.Name,
		Description: command.Description,
		Date:        command.Date,
		PartnerId:   command.PartnerId,
	}
	event, err := NewEvent(*props)
	return event, err
}

func (props *Event) AddSection(command AddSectionCommand) error {
	createEventSectionCommand := CreateEventSectionCommand{
		command.Name,
		command.Description,
		command.TotalSpot,
		command.Price,
	}
	eventSection, err := CreateEventSection(createEventSectionCommand)
	if eventSection != nil {
		props.sections = append(props.sections, *eventSection)
		props.totalSpots += eventSection.totalSpot
	}
	return err
}

func (props *Event) ChangeName(name string) error {
	newName, err := values_objects.NewName(name)
	if newName != nil {
		props.name = *newName
	}
	return err
}

func (props *Event) GetId() uuid.UUID {
	return props.id.GetValue()
}

func (props *Event) GetName() string {
	return props.name.GetValue()
}

func (props *Event) GetDescription() string {
	return props.description
}

func (props *Event) GetDate() time.Time {
	return props.date
}

func (props *Event) IsPublished() bool {
	return props.published
}

func (props *Event) GetTotalSpots() int {
	return props.totalSpots
}

func (props *Event) GetTotalSpotReserved() int {
	return props.totalSpotReserved
}

func (props *Event) GetPartnerId() uuid.UUID {
	return props.partnerId.GetValue()
}

func (props *Event) GetSections() []EventSection {
	return props.sections
}
