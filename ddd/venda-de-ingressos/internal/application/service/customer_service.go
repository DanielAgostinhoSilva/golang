package service

import (
	"venda-de-ingressos/internal/domain/entities"
	domain "venda-de-ingressos/internal/domain/repositories"
)

type RegisterDto struct {
	Name string
	CPF  string
}

type CustomerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomerService(customerAdapter domain.CustomerRepository) *CustomerService {
	return &CustomerService{customerAdapter}
}

func (props *CustomerService) List() ([]*entities.Customer, error) {
	return props.customerRepository.FindAll()
}

func (props *CustomerService) Register(input RegisterDto) (*entities.Customer, error) {
	customer, err := entities.CreateCustomer(entities.CreateCustomerCommand{
		Name: input.Name,
		Cpf:  input.CPF,
	})
	if err != nil {
		return nil, err
	}
	err = props.customerRepository.Save(*customer)
	if err != nil {
		return nil, err
	}
	return customer, err
}
