package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type PartnerSuitTest struct {
	suite.Suite
	props1               PartnerProps
	createPartnerCommand CreatePartnerCommand
	initEventCommand     InitEventCommand
}

func (suite *PartnerSuitTest) SetupTest() {
	id, _ := uuid.Parse("bb46df01-2924-4a0a-a15e-4160a3284c55")
	suite.props1 = PartnerProps{
		id,
		"Partner Name",
	}
	suite.createPartnerCommand = CreatePartnerCommand{
		"Partner Name X",
	}
	suite.initEventCommand = InitEventCommand{
		"Partner Name Y",
		"Partner Description Y",
		time.Date(2023, time.August, 18, 0, 0, 0, 0, time.UTC),
	}
}

func (suite *PartnerSuitTest) Test_deve_inicializar_um_partner() {
	partner, err := NewPartner(suite.props1)
	suite.Nil(err)
	suite.Equal(suite.props1.Id, partner.GetId())
	suite.Equal(suite.props1.Name, partner.GetName())
}

func (suite *PartnerSuitTest) Test_deve_criar_um_partner_valido() {
	partner, err := CreatePartner(suite.createPartnerCommand)
	suite.Nil(err)
	suite.Equal(suite.createPartnerCommand.Name, partner.GetName())
}

func (suite *PartnerSuitTest) Test_deve_iniciar_um_event_a_partir_de_um_partner() {
	partner, _ := CreatePartner(suite.createPartnerCommand)
	event, err := partner.InitEvent(suite.initEventCommand)
	suite.Nil(err)
	suite.Equal(suite.initEventCommand.Name, event.GetName())
	suite.Equal(suite.initEventCommand.Description, event.GetDescription())
	suite.Equal(suite.initEventCommand.Date, event.GetDate())
	suite.Equal(partner.GetId(), event.GetPartnerId())
}

func (suite *PartnerSuitTest) Test_deve_permitir_alterar_o_nome_do_partner() {
	partner, _ := CreatePartner(suite.createPartnerCommand)
	err := partner.ChangeName("Partner Name abc")
	suite.Nil(err)
	suite.Equal("Partner Name abc", partner.GetName())
}

func Test_Partner(t *testing.T) {
	suite.Run(t, new(PartnerSuitTest))
}
