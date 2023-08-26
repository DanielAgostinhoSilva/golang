package model

import "github.com/google/uuid"

type CustomerModel struct {
	ID   uuid.UUID
	CPF  string
	NAME string
}

func (CustomerModel) TableName() string {
	return "CUSTOMER"
}
