package domain

import (
	"github.com/google/uuid"
	"venda-de-ingressos/internal/domain/entities"
)

type CustomerRepository interface {
	Save(entity entities.Customer) error
	FindById(id uuid.UUID) (entities.Customer, error)
	FindAll() ([]entities.Customer, error)
	Delete(id uuid.UUID) error
}
