package domain

import (
	"github.com/google/uuid"
	"venda-de-ingressos/internal/domain/entities"
)

type PartnerRepository interface {
	Save(entity entities.Partner) error
	FindById(id uuid.UUID) (entities.Partner, error)
	FindAll() ([]entities.Partner, error)
	Delete(id uuid.UUID) error
}
