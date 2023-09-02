package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"venda-de-ingressos/internal/domain/entities"
	"venda-de-ingressos/internal/infra/db/mappers"
	"venda-de-ingressos/internal/infra/db/model"
)

type EventAdapter struct {
	DB     *gorm.DB
	mapper mappers.EventMapper
}

func NewEventAdapter(db *gorm.DB) *EventAdapter {
	eventSpotMapper := mappers.NewEventSpot()
	eventSectionMapper := mappers.NewEventSection(eventSpotMapper)
	eventMapper := mappers.NewEventMapper(eventSectionMapper)
	return &EventAdapter{db, eventMapper}
}

func (props *EventAdapter) Save(entity entities.Event) error {
	eventModel := props.mapper.ToModel(entity)
	return props.DB.Create(eventModel).Error
}

func (props *EventAdapter) FindById(id uuid.UUID) (*entities.Event, error) {
	eventModel, err := props.findOrFail(id)
	if err != nil {
		return nil, err
	}
	return props.mapper.ToDomain(*eventModel)
}

func (props *EventAdapter) FindAll() ([]*entities.Event, error) {
	var models []model.EventModel
	err := props.DB.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return props.mapper.ToCollectionDomain(models)
}

func (props *EventAdapter) Delete(id uuid.UUID) error {
	model, err := props.findOrFail(id)
	if err != nil {
		return err
	}
	return props.DB.Delete(model).Error
}

func (props *EventAdapter) findOrFail(id uuid.UUID) (*model.EventModel, error) {
	var eventModel model.EventModel
	//err := props.DB.Preload("Sections").First(&eventModel, "id = ?", id).Error
	err := props.DB.Model(&model.EventModel{}).Preload("Sections").First(&eventModel, "id = ?", id).Error
	return &eventModel, err
}
