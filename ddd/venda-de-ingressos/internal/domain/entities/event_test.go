package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EventSuitTest struct {
	suite.Suite
	props1             EventProps
	createEventCommand CreateEventCommand
	addSectionCommand  AddSectionCommand
}

func (suite *EventSuitTest) SetupTest() {
	id, _ := uuid.Parse("c381befe-e79e-498e-a237-fe5a8b0539c2")
	partnerId, _ := uuid.Parse("1226f73f-ec21-4f23-a75b-9eb63caac5bb")
	suite.props1 = EventProps{
		Id:                id,
		Name:              "Event Name A",
		Description:       "Event Description A",
		Date:              time.Date(2023, time.August, 18, 0, 0, 0, 0, time.UTC),
		Published:         false,
		TotalSpots:        0,
		TotalSpotReserved: 0,
		PartnerId:         partnerId,
		Sections:          []EventSection{},
	}
	suite.createEventCommand = CreateEventCommand{
		Name:        "Event Name A",
		Description: "Event Description A",
		Date:        time.Date(2023, time.August, 18, 0, 0, 0, 0, time.UTC),
		PartnerId:   partnerId,
	}
	suite.addSectionCommand = AddSectionCommand{
		Name:        "Event Name A",
		Description: "Event Description A",
		TotalSpot:   10,
		Price:       100.00,
	}
}

func (suite *EventSuitTest) Test_deve_inicializar_um_event() {
	event, err := NewEvent(suite.props1)
	suite.Nil(err)
	suite.Equal(suite.props1.Name, event.GetName())
	suite.Equal(suite.props1.Description, event.GetDescription())
	suite.Equal(suite.props1.Date, event.GetDate())
	suite.False(event.IsPublished())
	suite.Equal(suite.props1.TotalSpots, event.GetTotalSpots())
	suite.Equal(suite.props1.TotalSpotReserved, event.GetTotalSpotReserved())
	suite.Equal(suite.props1.PartnerId, event.GetPartnerId())
	suite.Len(event.GetSections(), 0)
}

func (suite *EventSuitTest) Test_deve_criar_um_evento_valido() {
	event, err := CreateEvent(suite.createEventCommand)
	suite.Nil(err)
	suite.Equal(suite.createEventCommand.Name, event.GetName())
	suite.Equal(suite.createEventCommand.Description, event.GetDescription())
	suite.Equal(suite.createEventCommand.Date, event.GetDate())
	suite.Equal(suite.createEventCommand.PartnerId, event.GetPartnerId())
}

func (suite *EventSuitTest) Test_deve_permitir_adicionar_uma_event_desction() {
	event, _ := CreateEvent(suite.createEventCommand)
	err := event.AddSection(suite.addSectionCommand)
	suite.Nil(err)
	suite.Equal(suite.addSectionCommand.TotalSpot, event.totalSpots)
	suite.Len(event.GetSections(), 1)
}

func (suite *EventSuitTest) Test_deve_permitir_publicar_um_event() {
	event, _ := CreateEvent(suite.createEventCommand)
	event.Publish()
	suite.True(event.IsPublished())
}

func (suite *EventSuitTest) Test_deve_permitir_despublicar_um_event() {
	event, _ := CreateEvent(suite.createEventCommand)
	event.UnPublish()
	suite.False(event.IsPublished())
}

func Test_EventSuitTest(t *testing.T) {
	suite.Run(t, new(EventSuitTest))
}
