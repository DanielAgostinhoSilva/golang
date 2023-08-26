package entities

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"venda-de-ingressos/pkg/domain/values-objects"
)

var (
	ErrCpfIsRequired = errors.New("Cpfcpf is required")
)

type CreateCustomerCommand struct {
	Name string
	Cpf  string
}

type CustomerProps struct {
	Id   uuid.UUID
	Cpf  string
	Name string
}

type Customer struct {
	id   uuid.UUID
	cpf  values_objects.Cpf
	name values_objects.Name
}

func NewCustomer(props CustomerProps) (*Customer, error) {
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
	customer, err := NewCustomer(CustomerProps{
		uuid.New(),
		command.Cpf,
		command.Name,
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
	customerProps := &CustomerProps{props.id, props.cpf.GetValue(), props.name.GetValue()}
	jsonProps, _ := json.Marshal(customerProps)
	return string(jsonProps)
}
