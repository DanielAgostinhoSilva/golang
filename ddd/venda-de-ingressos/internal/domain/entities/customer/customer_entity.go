package customer

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"venda-de-ingressos/pkg/domain/values-objects"
)

var (
	ErrCpfIsRequired = errors.New("cpf is required")
)

type CreateCustomerCommand struct {
	name string
	cpf  string
}

type Props struct {
	Id   uuid.UUID
	Cpf  string
	Name string
}

type Customer struct {
	id   uuid.UUID
	cpf  values_objects.Cpf
	name values_objects.Name
}

func NewCustomer(props Props) (*Customer, error) {
	name, err := values_objects.NewName(props.Name)
	if err != nil {
		return nil, err
	}

	cpf, err := values_objects.NewCpf(props.Cpf)
	if err != nil {
		return nil, err
	}

	return &Customer{
		id:   props.Id,
		cpf:  *cpf,
		name: *name,
	}, nil
}

func CreateCustomer(command CreateCustomerCommand) (*Customer, error) {
	customer, err := NewCustomer(Props{
		uuid.New(),
		command.cpf,
		command.name,
	})
	return customer, err
}

func (props *Customer) GetId() uuid.UUID {
	return props.id
}

func (props *Customer) GetName() *values_objects.Name {
	return &props.name
}

func (props *Customer) GetCpf() *values_objects.Cpf {
	return &props.cpf
}

func (props *Customer) Equals(customer Customer) bool {
	return props.id == customer.id
}

func (props *Customer) ToJson() string {
	customerProps := &Props{props.id, props.cpf.GetValue(), props.name.GetValue()}
	jsonProps, _ := json.Marshal(customerProps)
	return string(jsonProps)
}
