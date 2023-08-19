package events

import (
	"github.com/google/uuid"
	values_objects "venda-de-ingressos/pkg/domain/values-objects"
)

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

func (props *Partner) GetId() uuid.UUID {
	return props.id.GetValue()
}

func (props *Partner) GetName() string {
	return props.name.GetValue()
}
