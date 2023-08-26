package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EventSpotSuitTest struct {
	suite.Suite
	props1 EventSpotProps
}

func (suite *EventSpotSuitTest) SetupTest() {
	id, _ := uuid.Parse("2dc1b1cc-2468-485b-92f0-e7f0e366086e")
	suite.props1 = EventSpotProps{
		id,
		"Test location 1",
		false,
		false,
	}
}

func (suite *EventSpotSuitTest) Test_deve_inicilizar_um_EventSpot() {
	eventSpot, err := NewEventSpot(suite.props1)
	suite.Nil(err)
	suite.Equal(suite.props1.Id, eventSpot.GetId())
	suite.Equal(suite.props1.Location, eventSpot.GetLocation())
	suite.Equal(suite.props1.Reserved, eventSpot.IsReserved())
	suite.Equal(suite.props1.Published, eventSpot.IsPublished())
}

func (suite *EventSpotSuitTest) Test_deve_permitir_criar_um_event_spot_valido() {
	eventSpot, err := CreateEventSpot()
	suite.Nil(err)
	suite.Len(eventSpot.GetId(), 16)
	suite.Empty(eventSpot.GetLocation())
	suite.False(eventSpot.IsReserved())
	suite.False(eventSpot.IsPublished())
}

func (suite *EventSpotSuitTest) Test_deve_permitir_alterar_o_location() {
	eventSpot, _ := NewEventSpot(suite.props1)
	eventSpot.ChangeLocation("Location Test X")
	suite.Equal("Location Test X", eventSpot.GetLocation())
}

func (suite *EventSpotSuitTest) Test_deve_permitir_fazer_um_publish_no_EventSpot() {
	eventSpot, _ := NewEventSpot(suite.props1)
	eventSpot.Publish()
	suite.True(eventSpot.IsPublished())
}

func (suite *EventSpotSuitTest) Test_deve_permitir_fazer_um_UnPublish_no_EventSpot() {
	eventSpot, _ := NewEventSpot(suite.props1)
	eventSpot.UnPublish()
	suite.False(eventSpot.IsPublished())
}

func (suite *EventSpotSuitTest) Test_deve_exibir_um_json_quando_chamar_o_metodo_ToJson() {
	eventSpot, _ := NewEventSpot(suite.props1)
	suite.Equal(
		`{"ID":"2dc1b1cc-2468-485b-92f0-e7f0e366086e","Location":"Test location 1","Reserved":false,"Published":false}`,
		eventSpot.ToJson())
}

func Test_EventSpot(t *testing.T) {
	suite.Run(t, new(EventSpotSuitTest))
}
