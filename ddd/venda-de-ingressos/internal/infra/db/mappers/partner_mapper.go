package mappers

import (
	"venda-de-ingressos/internal/domain/entities"
	"venda-de-ingressos/internal/infra/db/model"
)

type PartnerMapper struct {
}

func NewPartnerMapper() PartnerMapper {
	return PartnerMapper{}
}

func (props *PartnerMapper) ToDomain(model model.PartnerModel) (*entities.Partner, error) {
	return entities.NewPartner(entities.PartnerProps{
		Id:   model.Id,
		Name: model.Name,
	})
}

func (props *PartnerMapper) ToCollectionDomain(models []model.PartnerModel) ([]*entities.Partner, error) {
	var partners []*entities.Partner
	for _, model := range models {
		partner, err := props.ToDomain(model)
		if err != nil {
			return nil, err
		}
		partners = append(partners, partner)
	}
	return partners, nil
}
