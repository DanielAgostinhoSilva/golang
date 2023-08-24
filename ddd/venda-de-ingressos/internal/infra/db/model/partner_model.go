package model

import "github.com/google/uuid"

type PartnerModel struct {
	Id   uuid.UUID
	Name string
}

func (PartnerModel) TableName() string {
	return "Partner"
}
