package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"venda-de-ingressos/internal/domain/entities"
	"venda-de-ingressos/internal/infra/db/model"
)

type PartnerAdapter struct {
	DB *gorm.DB
}

func NewPartnerAdapter(db *gorm.DB) *PartnerAdapter {
	return &PartnerAdapter{db}
}

func (props *PartnerAdapter) Save(partner entities.Partner) error {
	return props.DB.Create(&model.PartnerModel{
		Id:   partner.GetId(),
		Name: partner.GetName(),
	}).Error
}

func (props *PartnerAdapter) FindById(id uuid.UUID) (*entities.Partner, error) {
	var partnerModel model.PartnerModel
	err := props.DB.First(&partnerModel, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	partnerProps := entities.PartnerProps{
		Id:   partnerModel.Id,
		Name: partnerModel.Name,
	}
	return entities.NewPartner(partnerProps)
}
