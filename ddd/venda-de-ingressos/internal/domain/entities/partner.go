package entities

import (
	"github.com/google/uuid"
	"time"
	values_objects "venda-de-ingressos/pkg/domain/values-objects"
)

type InitEventCommand struct {
	Name        string
	Description string
	Date        time.Time
}

type CreatePartnerCommand struct {
	Name string
}

type PartnerProps struct {
	Id   uuid.UUID
	Name string
}

type Partner struct {
	id   PartnerId
	name values_objects.Name
}

func NewPartner(props PartnerProps) (*Partner, error) {
	name, err := values_objects.NewName(props.Name)
	if err != nil {
		return nil, err
	}
	return &Partner{
		*NewPartnerId(props.Id),
		*name,
	}, err
}

func CreatePartner(command CreatePartnerCommand) (*Partner, error) {
	props := PartnerProps{
		uuid.New(),
		command.Name,
	}
	return NewPartner(props)
}

func (props *Partner) InitEvent(command InitEventCommand) (*Event, error) {
	event, err := CreateEvent(CreateEventCommand{
		command.Name,
		command.Description,
		command.Date,
		props.id.GetValue(),
	})
	return event, err
}

func (props *Partner) ChangeName(name string) error {
	newName, err := values_objects.NewName(name)
	if newName != nil {
		props.name = *newName
	}
	return err
}

func (props *Partner) GetId() uuid.UUID {
	return props.id.GetValue()
}

func (props *Partner) GetName() string {
	return props.name.GetValue()
}
