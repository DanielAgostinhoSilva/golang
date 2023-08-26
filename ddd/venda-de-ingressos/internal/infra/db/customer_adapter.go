package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"venda-de-ingressos/internal/domain/entities"
	"venda-de-ingressos/internal/infra/db/mappers"
	"venda-de-ingressos/internal/infra/db/model"
)

type CustomerAdapter struct {
	DB     *gorm.DB
	mapper mappers.CustomerMapper
}

func NewCustomerAdapter(db *gorm.DB) *CustomerAdapter {
	return &CustomerAdapter{db, mappers.NewCustomerMapper()}
}

func (props *CustomerAdapter) Save(customer entities.Customer) error {
	return props.DB.Create(&model.CustomerModel{
		ID:   customer.GetId(),
		NAME: customer.GetName().GetValue(),
		CPF:  customer.GetCpf().GetValue(),
	}).Error
}

func (props *CustomerAdapter) FindById(id uuid.UUID) (*entities.Customer, error) {
	customerModel, err := props.findOrFail(id)
	if err != nil {
		return nil, err
	}
	return props.mapper.ToDomain(*customerModel)
}

func (props *CustomerAdapter) FindAll() ([]*entities.Customer, error) {
	var models []model.CustomerModel
	err := props.DB.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return props.mapper.ToCollectionDomain(models)
}

func (props *CustomerAdapter) Delete(id uuid.UUID) error {
	partner, err := props.findOrFail(id)
	if err != nil {
		return err
	}
	return props.DB.Delete(partner).Error
}

func (props *CustomerAdapter) findOrFail(id uuid.UUID) (*model.CustomerModel, error) {
	var customerModel model.CustomerModel
	err := props.DB.First(&customerModel, "id = ?", id).Error
	return &customerModel, err
}
