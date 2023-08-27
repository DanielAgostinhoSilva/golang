package model

import "github.com/google/uuid"

type EventSpotModel struct {
	id        uuid.UUID `gorm:"column:id"`
	location  string    `gorm:"column:location"`
	reserved  bool      `gorm:"column:reserved"`
	published bool      `gorm:"column:published"`
}

func (EventSpotModel) TableName() string {
	return "event_spot"
}
