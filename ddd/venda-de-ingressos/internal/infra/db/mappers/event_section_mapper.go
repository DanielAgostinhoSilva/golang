package mappers

import (
	"venda-de-ingressos/internal/domain/entities"
	"venda-de-ingressos/internal/infra/db/model"
)

type EventSectionMapper struct {
	EventSpotMapper EventSpotMapper
}

func NewEventSection(eventSpotMapper EventSpotMapper) EventSectionMapper {
	return EventSectionMapper{
		EventSpotMapper: eventSpotMapper,
	}
}

func (props *EventSectionMapper) ToDomain(model model.EventSectionModel) (*entities.EventSection, error) {
	spots, err := props.EventSpotMapper.ToCollectionDomain(model.Spots)
	if err != nil {
		return nil, err
	}

	return entities.NewEventSection(entities.EventSectionProps{
		Id:                model.Id,
		Name:              model.Name,
		Description:       model.Description,
		Published:         model.Published,
		TotalSpot:         model.TotalSpot,
		TotalSpotReserved: model.TotalSpotReserved,
		Price:             model.Price,
		Spots:             spots,
	})
}

func (props *EventSectionMapper) ToModel(entity entities.EventSection) *model.EventSectionModel {
	return &model.EventSectionModel{
		Id:                entity.GetId(),
		Name:              entity.GetName(),
		Description:       entity.GetDescription(),
		Published:         entity.IsPublished(),
		TotalSpot:         entity.GetTotalSpot(),
		TotalSpotReserved: entity.GetTotalSpotReserved(),
		Price:             entity.GetPrice(),
		Spots:             props.EventSpotMapper.ToCollectionModel(entity.GetSpots()),
	}
}

func (props *EventSectionMapper) ToCollectionDomain(models []model.EventSectionModel) ([]entities.EventSection, error) {
	var sections []entities.EventSection
	for _, model := range models {
		section, err := props.ToDomain(model)
		if err != nil {
			return nil, err
		}
		sections = append(sections, *section)
	}
	return sections, nil
}

func (props *EventSectionMapper) ToCollectionModel(entities []entities.EventSection) []model.EventSectionModel {
	var models []model.EventSectionModel
	for _, entity := range entities {
		model := props.ToModel(entity)
		models = append(models, *model)
	}
	return models
}
