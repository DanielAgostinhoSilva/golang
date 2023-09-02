package model

import (
	"github.com/google/uuid"
)

type EventSectionModel struct {
	Id                uuid.UUID `gorm:"primarykey"`
	Name              string
	Description       string
	Published         bool
	TotalSpot         int
	TotalSpotReserved int
	Price             float64
	Spots             []EventSpotModel `gorm:"foreignKey:EventSectionID"`
	EventId           uuid.UUID        `gorm:"column:event_id"`
}

func (EventSectionModel) TableName() string {
	return "event_section"
}
