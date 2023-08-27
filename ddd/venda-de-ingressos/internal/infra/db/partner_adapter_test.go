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
	db                   *gorm.DB
	createPartnerCommand entities.CreatePartnerCommand
	env                  *configs.EnvConfig
}

func (suite *PartnerAdapterSuiteTest) SetupSuite() {
	suite.env = configs.LoadEnvConfig("./../../../cmd/server/test.env")
	suite.db = configs.LoadDataBase(*suite.env)
	configs.LoadMigrationUp(*suite.env)
}

func (suite *PartnerAdapterSuiteTest) SetupTest() {
	suite.createPartnerCommand = entities.CreatePartnerCommand{
		Name: "Partner Name",
	}
}

func (suite *PartnerAdapterSuiteTest) TearDownTest() {
	suite.db.Table("partner").Where("id is not null").Delete(nil)
}

func (suite *PartnerAdapterSuiteTest) TearDownSuite() {
	configs.LoadMigrationDown(*suite.env)
}

func (suite *PartnerAdapterSuiteTest) Test_deve_persistir_um_partner_no_banco_de_dados() {
	adapter := NewPartnerAdapter(suite.db)
	partner, err := entities.CreatePartner(suite.createPartnerCommand)
	suite.Nil(err)
	err = adapter.Save(*partner)
	suite.Nil(err)
}

func (suite *PartnerAdapterSuiteTest) Test_deve_buscar_um_partner_por_id() {
	adapter := NewPartnerAdapter(suite.db)
	partner, err := entities.CreatePartner(suite.createPartnerCommand)
	err = adapter.Save(*partner)
	suite.Nil(err)

	partnerFound, err := adapter.FindById(partner.GetId())
	suite.Nil(err)
	suite.Equal(partner.GetId(), partnerFound.GetId())
}

func (suite *PartnerAdapterSuiteTest) Test_deve_falhar_quando_nao_conseguir_encontrar_um_partner() {
	adapter := NewPartnerAdapter(suite.db)
	partnerFound, err := adapter.FindById(uuid.New())
	suite.Nil(partnerFound)
	suite.Error(err)
}

func (suite *PartnerAdapterSuiteTest) Test_deve_buscar_todos_os_partners_no_banco_de_dados() {
	adapter := NewPartnerAdapter(suite.db)
	partner, err := entities.CreatePartner(suite.createPartnerCommand)
	suite.Nil(err)

	err = adapter.Save(*partner)
	suite.Nil(err)

	partners, err := adapter.FindAll()
	suite.Nil(err)
	suite.Len(partners, 1)
}

func (suite *PartnerAdapterSuiteTest) Test_deve_remover_um_partner_no_banco_de_dados() {
	adapter := NewPartnerAdapter(suite.db)
	partner, err := entities.CreatePartner(suite.createPartnerCommand)
	suite.Nil(err)

	err = adapter.Save(*partner)
	suite.Nil(err)

	err = adapter.Delete(partner.GetId())
	suite.Nil(err)
}

func (suite *PartnerAdapterSuiteTest) Test_deve_falhar_ao_remover_um_partner_que_nao_existe() {
	adapter := NewPartnerAdapter(suite.db)
	err := adapter.Delete(uuid.New())
	suite.Error(err)
}

func Test_PartnerAdapterSuite(t *testing.T) {
	suite.Run(t, new(PartnerAdapterSuiteTest))
}
