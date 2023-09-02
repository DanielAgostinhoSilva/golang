package model

import (
	"github.com/google/uuid"
	"time"
)

type EventModel struct {
	Id                 uuid.UUID `gorm:"primarykey"`
	Name               string
	Description        string
	Date               time.Time
	Published          bool
	TotalSpots         int
	TotalSpotsReserved int
	PartnerId          uuid.UUID
	Sections           []EventSectionModel `gorm:"foreignKey:EventId"`
}

func (EventModel) TableName() string {
	return "event"
}
