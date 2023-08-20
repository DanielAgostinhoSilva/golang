package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EventSectionSuitTest struct {
	suite.Suite
	props1   EventSectionProps
	command1 CreateEventSectionCommand
}

func (suite *EventSectionSuitTest) SetupTest() {
	id, _ := uuid.Parse("2dc1b1cc-2468-485b-92f0-e7f0e366086e")
	suite.props1 = EventSectionProps{
		id,
		"Test location 1",
		"Test Description",
		false,
		0,
		0,
		0.0,
		[]EventSpot{},
	}
	suite.command1 = CreateEventSectionCommand{
		"Test location 1",
		"Test Description",
		10,
		100.0,
	}
}

func (suite *EventSectionSuitTest) Test_deve_incializar_um_EventSection() {
	eventSection, err := NewEventSection(suite.props1)
	suite.Nil(err)
	suite.Equal(suite.props1.Id, eventSection.GetId())
	suite.Equal(suite.props1.Name, eventSection.GetName())
	suite.Equal(suite.props1.Description, eventSection.GetDescription())
	suite.Equal(suite.props1.Published, eventSection.IsPublished())
	suite.Equal(suite.props1.TotalSpot, eventSection.GetTotalSpot())
	suite.Equal(suite.props1.TotalSpotReserved, eventSection.GetTotalSpotReserved())
	suite.Equal(suite.props1.Price, eventSection.GetPrice())
	suite.Equal(suite.props1.Spots, eventSection.GetSpots())
}

func (suite *EventSectionSuitTest) Test_deve_permitir_criar_um_EventSection() {
	eventSection, err := CreateEventSection(suite.command1)
	suite.Nil(err)
	suite.Equal(suite.command1.Name, eventSection.GetName())
	suite.Equal(suite.command1.Description, eventSection.GetDescription())
	suite.Equal(suite.command1.TotalSpot, eventSection.GetTotalSpot())
	suite.Equal(suite.command1.Price, eventSection.GetPrice())
	suite.Len(eventSection.spots, suite.command1.TotalSpot)
}

func (suite *EventSectionSuitTest) Test_deve_permitir_alterar_o_name() {
	eventSection, _ := CreateEventSection(suite.command1)
	err := eventSection.ChangeName("Test X")
	suite.Nil(err)
	suite.Equal("Test X", eventSection.GetName())
}

func (suite *EventSectionSuitTest) Test_deve_permitir_alterar_o_description() {
	eventSection, _ := CreateEventSection(suite.command1)
	eventSection.ChangeDescription("Test Description X")
	suite.Equal("Test Description X", eventSection.GetDescription())
}

func (suite *EventSectionSuitTest) Test_deve_permitir_alterar_o_price() {
	eventSection, _ := CreateEventSection(suite.command1)
	err := eventSection.ChangePrice(200.00)
	suite.Nil(err)
	suite.Equal(200.0, eventSection.GetPrice())
}

func (suite *EventSectionSuitTest) Test_deve_lancar_um_erro_quando_o_price_for_um_valor_invalido() {
	eventSection, _ := CreateEventSection(suite.command1)
	err := eventSection.ChangePrice(-200.00)
	suite.Error(err)
}

func (suite *EventSectionSuitTest) Test_deve_permitir_publicar_um_event_section() {
	eventSection, _ := CreateEventSection(suite.command1)
	eventSection.Publish()
	suite.True(eventSection.IsPublished())
}

func (suite *EventSectionSuitTest) Test_deve_permitir_despublicar_um_event_section() {
	eventSection, _ := CreateEventSection(suite.command1)
	eventSection.UnPublish()
	suite.False(eventSection.IsPublished())
}

func (suite *EventSectionSuitTest) Test_deve_publicar_todos_os_event_spot() {
	eventSection, _ := CreateEventSection(suite.command1)
	eventSection.PublishAll()
	suite.True(eventSection.IsPublished())
	for _, spot := range eventSection.GetSpots() {
		suite.True(spot.IsPublished())
	}
}

func Test_EventSection(t *testing.T) {
	suite.Run(t, new(EventSectionSuitTest))
}
