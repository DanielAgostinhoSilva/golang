package model

import (
	"github.com/google/uuid"
)

type CustomerModel struct {
	Id   uuid.UUID `gorm:"primarykey"`
	Cpf  string
	Name string
}

func (CustomerModel) TableName() string {
	return "customer"
}
