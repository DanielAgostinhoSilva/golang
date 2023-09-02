package domain

import (
	"github.com/google/uuid"
	"venda-de-ingressos/internal/domain/entities"
)

type EventRepository interface {
	Save(entity entities.Event) error
	FindById(id uuid.UUID) (entities.Event, error)
	FindAll() ([]entities.Event, error)
	Delete(id uuid.UUID) error
}
