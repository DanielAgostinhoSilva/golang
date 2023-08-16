package customer

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrIdIsRequired   = errors.New("id is required")
	ErrInvalidId      = errors.New("invalid id")
	ErrNameIsRequired = errors.New("name is required")
	ErrCpfIsRequired  = errors.New("cpf is required")
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
	cpf  string
	name string
}

func NewCustomer(props Props) (*Customer, error) {
	return &Customer{
		id:   props.Id,
		cpf:  props.Cpf,
		name: props.Name,
	}, nil
}

func CreateCustomer(command CreateCustomerCommand) (*Customer, error) {
	if command.name == "" {
		return nil, ErrNameIsRequired
	}

	if command.cpf == "" {
		return nil, ErrCpfIsRequired
	}

	return &Customer{
		id:   uuid.New(),
		name: command.name,
		cpf:  command.cpf,
	}, nil
}

func (props *Customer) GetId() uuid.UUID {
	return props.id
}

func (props *Customer) GetName() string {
	return props.name
}

func (props *Customer) GetCpf() string {
	return props.cpf
}

func (props *Customer) Equals(customer Customer) bool {
	return props.id == customer.id
}

func (props *Customer) ToJson() string {
	customerProps := &Props{props.id, props.cpf, props.name}
	jsonProps, _ := json.Marshal(customerProps)
	return string(jsonProps)
}
