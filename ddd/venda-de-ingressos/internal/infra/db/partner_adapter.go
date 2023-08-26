package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"venda-de-ingressos/internal/domain/entities"
	"venda-de-ingressos/internal/infra/db/mappers"
	"venda-de-ingressos/internal/infra/db/model"
)

type PartnerAdapter struct {
	DB     *gorm.DB
	mapper mappers.PartnerMapper
}

func NewPartnerAdapter(db *gorm.DB) *PartnerAdapter {
	return &PartnerAdapter{db, mappers.NewPartnerMapper()}
}

func (props *PartnerAdapter) Save(partner entities.Partner) error {
	return props.DB.Create(&model.PartnerModel{
		Id:   partner.GetId(),
		Name: partner.GetName(),
	}).Error
}

func (props *PartnerAdapter) FindById(id uuid.UUID) (*entities.Partner, error) {
	partnerModel, err := props.findOrFail(id)
	if err != nil {
		return nil, err
	}
	return props.mapper.ToDomain(*partnerModel)
}

func (props *PartnerAdapter) FindAll() ([]*entities.Partner, error) {
	var models []model.PartnerModel
	err := props.DB.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return props.mapper.ToCollectionDomain(models)
}

func (props *PartnerAdapter) Delete(id uuid.UUID) error {
	partner, err := props.findOrFail(id)
	if err != nil {
		return err
	}
	return props.DB.Delete(partner).Error
}

func (props *PartnerAdapter) findOrFail(id uuid.UUID) (*model.PartnerModel, error) {
	var partnerModel model.PartnerModel
	err := props.DB.First(&partnerModel, "id = ?", id).Error
	return &partnerModel, err
}
