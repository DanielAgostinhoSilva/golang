package db

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"venda-de-ingressos/configs"
	"venda-de-ingressos/internal/domain/entities"
)

type PartnerAdapterSuiteTest struct {
	suite.Suite
	db           *gorm.DB
	partnerProps entities.PartnerProps
}

func (suite *PartnerAdapterSuiteTest) SetupTest() {
	id, _ := uuid.Parse("bb46df01-2924-4a0a-a15e-4160a3284c55")
	suite.partnerProps = entities.PartnerProps{
		Id:   id,
		Name: "Partner Name",
	}
	suite.db = configs.LoadSqlite("./../../../test.db")
	configs.LoadMigration("./../../../test.db", "sqlite3", "./migration")
}

func (suite *PartnerAdapterSuiteTest) TearDownSuite() {
	configs.LoadMigrationWithCommand("./../../../test.db", "sqlite3", "./migration", "down")
}

func (suite *PartnerAdapterSuiteTest) Test_deve_persistir_um_partner_no_banco_de_dados() {
	partner, err := entities.NewPartner(suite.partnerProps)
	suite.Nil(err)

	adapter := NewPartnerAdapter(suite.db)
	err = adapter.Save(*partner)
	suite.Nil(err)

	partnerFound, err := adapter.FindById(partner.GetId())
	suite.Nil(err)
	suite.Equal(partner.GetId(), partnerFound.GetId())
}

func Test_PartnerAdapterSuite(t *testing.T) {
	suite.Run(t, new(PartnerAdapterSuiteTest))
}
