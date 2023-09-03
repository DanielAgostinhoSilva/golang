package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"venda-de-ingressos/internal/domain/entities"
)

type CustomerServerSuiteTest struct {
	suite.Suite
	//mock.Mock
	customer        *entities.Customer
	customerAdapter *MockCustomerRepository
}

func (suite *CustomerServerSuiteTest) SetupTest() {
	id, _ := uuid.Parse("56b0fe0f-d186-49e9-9d38-7503d4a2a485")
	customer, _ := entities.NewCustomer(entities.CustomerProps{
		Id:   id,
		Cpf:  "188.539.920-00",
		Name: "CustomerA",
	})
	suite.customer = customer
	suite.customerAdapter = new(MockCustomerRepository)
}

func (suite *CustomerServerSuiteTest) Test_deve_listar_todo_os_customer() {
	suite.customerAdapter.On("FindAll").Return([]*entities.Customer{suite.customer}, nil)
	customerService := NewCustomerService(suite.customerAdapter)
	customers, err := customerService.List()
	suite.Nil(err)
	suite.Len(customers, 1)
	suite.Equal(suite.customer, customers[0])
}

func (suite *CustomerServerSuiteTest) Test_deve_retornar_um_erro_quando_tentar_listar_todos_os_customers() {
	suite.customerAdapter.On("FindAll").Return([]*entities.Customer{}, errors.New("test"))
	customerService := NewCustomerService(suite.customerAdapter)
	customers, err := customerService.List()
	suite.Empty(customers)
	suite.Error(err)
}

func (suite *CustomerServerSuiteTest) Test_deve_criar_um_customer() {
	suite.customerAdapter.On("Save", mock.AnythingOfType("Customer")).Return(nil)
	customerService := NewCustomerService(suite.customerAdapter)
	customer, err := customerService.Register(RegisterDto{
		suite.customer.GetName().GetValue(),
		suite.customer.GetCpf().GetValue(),
	})
	suite.Nil(err)
	suite.Equal(suite.customer.GetName(), customer.GetName())
	suite.Equal(suite.customer.GetCpf(), customer.GetCpf())
}

func (suite *CustomerServerSuiteTest) Test_deve_retonar_um_erro_ao_tentar_criar_um_customer() {
	customerAdapter := new(MockCustomerRepository)
	customerAdapter.On("Save", mock.AnythingOfType("Customer")).Return(errors.New("test"))
	customerService := NewCustomerService(customerAdapter)
	customer, err := customerService.Register(RegisterDto{
		suite.customer.GetName().GetValue(),
		suite.customer.GetCpf().GetValue(),
	})
	suite.Nil(customer)
	suite.Error(err)
}

func Test_CustomerServerSuite(t *testing.T) {
	suite.Run(t, new(CustomerServerSuiteTest))
}

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) Save(entity entities.Customer) error {
	args := m.Called(entity)
	return args.Error(0)
}

func (m *MockCustomerRepository) FindById(id uuid.UUID) (*entities.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Customer), args.Error(1)
}

func (m *MockCustomerRepository) FindAll() ([]*entities.Customer, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
