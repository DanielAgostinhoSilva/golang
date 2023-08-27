package model

import "github.com/google/uuid"

type PartnerModel struct {
	Id   uuid.UUID `gorm:"column:id"`
	Name string    `gorm:"column:name"`
}

func (PartnerModel) TableName() string {
	return "partner"
}
