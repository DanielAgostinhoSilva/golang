package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"venda-de-ingressos/internal/domain/entities"
)

type PartnerModel struct {
	Id   uuid.UUID
	Name string
}

type PartnerAdapter struct {
	DB *gorm.DB
}

func NewPartnerAdapter(db *gorm.DB) *PartnerAdapter {
	return &PartnerAdapter{db}
}

func (props *PartnerAdapter) Save(partner entities.Partner) error {
	return props.DB.Create(&PartnerModel{
		partner.GetId(),
		partner.GetName(),
	}).Error
}
