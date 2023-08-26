package model

import "github.com/google/uuid"

type PartnerModel struct {
	ID   uuid.UUID
	NAME string
}

func (PartnerModel) TableName() string {
	return "PARTNER"
}
