package customer

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	props := &CustomerProps{
		Id:   uuid.New(),
		Cpf:  "290.658.340-52",
		Name: "Test A",
	}
	customer, err := NewCustomer(*props)
	assert.Nil(t, err)
	assert.Equal(t, props.Id, customer.GetId())
	assert.Equal(t, props.Name, customer.GetName())
	assert.Equal(t, props.Cpf, customer.GetCpf())
}

func TestCreateCustomer(t *testing.T) {
	command := &CreateCustomerCommand{"Test A", "290.658.340-52"}
	customer, err := CreateCustomer(*command)
	assert.Nil(t, err)
	assert.Equal(t, command.name, customer.GetName())
	assert.Equal(t, command.cpf, customer.GetCpf())
}
