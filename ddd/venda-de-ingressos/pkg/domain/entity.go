package domain

import (
	"github.com/google/uuid"
)

type Entity interface {
	GetId() uuid.UUID
	Equals(entity Entity) bool
	ToJson() string
}
