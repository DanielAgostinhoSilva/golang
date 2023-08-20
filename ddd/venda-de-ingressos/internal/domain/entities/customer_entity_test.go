package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CustomerTestSuit struct {
	suite.Suite
	customerProps1         Props
	customerProps2         Props
	createCustomerCommand1 CreateCustomerCommand
}

func (suite *CustomerTestSuit) SetupTest() {
	id, _ := uuid.Parse("5b2c9913-f782-4dab-a950-99650497c115")
	suite.customerProps1 = Props{
		Id:   id,
		Cpf:  "29065834052",
		Name: "Test A",
	}
	id, _ = uuid.Parse("bf913ce3-2f8c-446a-b96f-7966486bffa4")
	suite.customerProps1 = Props{
		Id:   id,
		Cpf:  "41152382080",
		Name: "Test B",
	}
	suite.createCustomerCommand1 = CreateCustomerCommand{
		name: suite.customerProps1.Name,
		cpf:  suite.customerProps1.Cpf,
	}
}

func (suite *CustomerTestSuit) Test_deve_inicializar_um_Customer() {
	customer, err := NewCustomer(suite.customerProps1)
	suite.Nil(err)
	suite.Equal(suite.customerProps1.Id, customer.GetId())
	suite.Equal(suite.customerProps1.Name, customer.GetName().GetValue())
	suite.Equal(suite.customerProps1.Cpf, customer.GetCpf().GetValue())
}

func (suite *CustomerTestSuit) Test_deve_criar_um_Customer_quando_receber_um_comando_CreateCustomerCommand() {
	customer, err := CreateCustomer(suite.createCustomerCommand1)
	suite.Nil(err)
	suite.Equal(suite.createCustomerCommand1.name, customer.GetName().GetValue())
	suite.Equal(suite.createCustomerCommand1.cpf, customer.GetCpf().GetValue())
}

func (suite *CustomerTestSuit) Test_deve_verificar_se_um_Customer_e_igual_ao_outro() {
	customerA, _ := NewCustomer(suite.customerProps1)
	customerB, _ := CreateCustomer(suite.createCustomerCommand1)
	customerC, _ := NewCustomer(suite.customerProps1)
	suite.False(customerA.Equals(*customerB))
	suite.True(customerA.Equals(*customerC))
}

func Test_Customer(t *testing.T) {
	suite.Run(t, new(CustomerTestSuit))
}
