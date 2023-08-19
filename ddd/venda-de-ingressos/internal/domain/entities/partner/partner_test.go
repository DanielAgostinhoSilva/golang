package events

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PartnerSuitTest struct {
	suite.Suite
	props1               PartnerProps
	createPartnerCommand CreatePartnerCommand
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

func Test_EventSection(t *testing.T) {
	suite.Run(t, new(PartnerSuitTest))
}
