package model

import (
	"github.com/google/uuid"
)

type EventSpotModel struct {
	Id             uuid.UUID `gorm:"primarykey"`
	Location       string
	Reserved       bool
	Published      bool
	EventSectionID uuid.UUID `gorm:"column:event_section_id"`
}

func (EventSpotModel) TableName() string {
	return "event_spot"
}
