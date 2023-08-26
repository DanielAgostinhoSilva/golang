package mappers

import (
	"venda-de-ingressos/internal/domain/entities"
	"venda-de-ingressos/internal/infra/db/model"
)

type CustomerMapper struct {
}

func NewCustomerMapper() CustomerMapper {
	return CustomerMapper{}
}

func (props *CustomerMapper) ToDomain(model model.CustomerModel) (*entities.Customer, error) {
	return entities.NewCustomer(entities.CustomerProps{
		Id:   model.ID,
		Cpf:  model.CPF,
		Name: model.NAME,
	})
}

func (props *CustomerMapper) ToCollectionDomain(models []model.CustomerModel) ([]*entities.Customer, error) {
	var customers []*entities.Customer
	for _, model := range models {
		customer, err := props.ToDomain(model)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
