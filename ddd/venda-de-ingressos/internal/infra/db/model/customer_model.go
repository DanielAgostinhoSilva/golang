package model

import "github.com/google/uuid"

type CustomerModel struct {
	Id   uuid.UUID `gorm:"column:id"`
	Cpf  string    `gorm:"column:cpf"`
	Name string    `gorm:"column:name"`
}

func (CustomerModel) TableName() string {
	return "customer"
}
