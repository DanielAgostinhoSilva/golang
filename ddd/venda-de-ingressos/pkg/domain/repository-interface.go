package domain

import "github.com/google/uuid"

type Repository interface {
	Save(entity interface{}) (interface{}, error)
	FindById(id uuid.UUID) (interface{}, error)
	FindAll() ([]interface{}, error)
	Delete(entity interface{}) (interface{}, error)
}
