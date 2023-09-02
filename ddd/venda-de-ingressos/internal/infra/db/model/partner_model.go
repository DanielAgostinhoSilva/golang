package model

import (
	"github.com/google/uuid"
)

type PartnerModel struct {
	Id   uuid.UUID `gorm:"primarykey"`
	Name string
}

func (PartnerModel) TableName() string {
	return "partner"
}
