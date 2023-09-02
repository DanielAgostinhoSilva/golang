package mappers

import (
	"venda-de-ingressos/internal/domain/entities"
	"venda-de-ingressos/internal/infra/db/model"
)

type EventSpotMapper struct {
}

func NewEventSpot() EventSpotMapper {
	return EventSpotMapper{}
}

func (props *EventSpotMapper) ToDomain(model model.EventSpotModel) (*entities.EventSpot, error) {
	return entities.NewEventSpot(entities.EventSpotProps{
		Id:        model.Id,
		Location:  model.Location,
		Reserved:  model.Reserved,
		Published: model.Published,
	})
}

func (props *EventSpotMapper) ToModel(entity entities.EventSpot) *model.EventSpotModel {
	return &model.EventSpotModel{
		Id:        entity.GetId(),
		Location:  entity.GetLocation(),
		Reserved:  entity.IsReserved(),
		Published: entity.IsPublished(),
	}
}

func (props *EventSpotMapper) ToCollectionDomain(models []model.EventSpotModel) ([]entities.EventSpot, error) {
	var partners []entities.EventSpot
	for _, model := range models {
		partner, err := props.ToDomain(model)
		if err != nil {
			return nil, err
		}
		partners = append(partners, *partner)
	}
	return partners, nil
}

func (props *EventSpotMapper) ToCollectionModel(entities []entities.EventSpot) []model.EventSpotModel {
	var models []model.EventSpotModel
	for _, entity := range entities {
		model := props.ToModel(entity)
		models = append(models, *model)
	}
	return models
}
