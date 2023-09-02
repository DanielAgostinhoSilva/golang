package mappers

import (
	"venda-de-ingressos/internal/domain/entities"
	"venda-de-ingressos/internal/infra/db/model"
)

type EventMapper struct {
	EventSectionMapper EventSectionMapper
}

func NewEventMapper(eventSectionMapper EventSectionMapper) EventMapper {
	return EventMapper{
		eventSectionMapper,
	}
}

func (props *EventMapper) ToDomain(model model.EventModel) (*entities.Event, error) {
	sections, err := props.EventSectionMapper.ToCollectionDomain(model.Sections)
	if err != nil {
		return nil, err
	}
	return entities.NewEvent(entities.EventProps{
		Id:                model.Id,
		Name:              model.Name,
		Description:       model.Description,
		Date:              model.Date,
		Published:         model.Published,
		TotalSpots:        model.TotalSpots,
		TotalSpotReserved: model.TotalSpotsReserved,
		PartnerId:         model.Id,
		Sections:          sections,
	})
}

func (props *EventMapper) ToModel(entity entities.Event) *model.EventModel {
	return &model.EventModel{
		Id:                 entity.GetId(),
		Name:               entity.GetName(),
		Description:        entity.GetDescription(),
		Date:               entity.GetDate(),
		Published:          entity.IsPublished(),
		TotalSpots:         entity.GetTotalSpots(),
		TotalSpotsReserved: entity.GetTotalSpotReserved(),
		PartnerId:          entity.GetPartnerId(),
		Sections:           props.EventSectionMapper.ToCollectionModel(entity.GetSections()),
	}
}

func (props *EventMapper) ToCollectionDomain(models []model.EventModel) ([]*entities.Event, error) {
	var events []*entities.Event
	for _, model := range models {
		event, err := props.ToDomain(model)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (props *EventMapper) ToCollectionModel(entity []entities.Event) []model.EventModel {
	var models []model.EventModel
	for _, entity := range entity {
		model := props.ToModel(entity)
		models = append(models, *model)
	}
	return models
}
