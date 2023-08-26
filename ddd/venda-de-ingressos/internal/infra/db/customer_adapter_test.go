package db

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"venda-de-ingressos/configs"
	"venda-de-ingressos/internal/domain/entities"
)

type CustomerAdapterSuiteTest struct {
	suite.Suite
	db                    *gorm.DB
	createCustomerCommand entities.CreateCustomerCommand
}

func (suite *CustomerAdapterSuiteTest) SetupSuite() {
	suite.db = configs.LoadSqlite("./../../../test.db")
	configs.LoadMigration("./../../../test.db", "sqlite3", "./migration")
}

func (suite *CustomerAdapterSuiteTest) SetupTest() {
	suite.createCustomerCommand = entities.CreateCustomerCommand{
		Name: "Partner NAME",
		Cpf:  "17711109024",
	}
}

func (suite *CustomerAdapterSuiteTest) TearDownTest() {
	suite.db.Table("CUSTOMER").Where("id is not null").Delete(nil)
}

func (suite *CustomerAdapterSuiteTest) TearDownSuite() {
	configs.LoadMigrationWithCommand("./../../../test.db", "sqlite3", "./migration", "down")
}

func (suite *CustomerAdapterSuiteTest) Test_deve_persistir_um_customer_no_banco_de_dados() {
	adapter := NewCustomerAdapter(suite.db)
	customer, err := entities.CreateCustomer(suite.createCustomerCommand)
	suite.Nil(err)
	err = adapter.Save(*customer)
	suite.Nil(err)
}

func (suite *CustomerAdapterSuiteTest) Test_deve_buscar_um_customer_por_id() {
	adapter := NewCustomerAdapter(suite.db)
	customer, err := entities.CreateCustomer(suite.createCustomerCommand)
	suite.Nil(err)

	err = adapter.Save(*customer)
	suite.Nil(err)

	customerFound, err := adapter.FindById(customer.GetId())
	suite.Nil(err)
	suite.Equal(customerFound.GetId(), customer.GetId())
}

func (suite *CustomerAdapterSuiteTest) Test_deve_retornar_um_erro_quando_buscar_um_customer_que_nao_existe() {
	adapter := NewCustomerAdapter(suite.db)
	customer, err := adapter.FindById(uuid.New())
	suite.Nil(customer)
	suite.Error(err)
}

func (suite *CustomerAdapterSuiteTest) Test_deve_buscar_todos_customer_no_banco_de_dados() {
	adapter := NewCustomerAdapter(suite.db)
	customer, err := entities.CreateCustomer(suite.createCustomerCommand)
	suite.Nil(err)

	err = adapter.Save(*customer)
	suite.Nil(err)

	customers, err := adapter.FindAll()
	suite.Nil(err)
	suite.Len(customers, 1)
}

func (suite *CustomerAdapterSuiteTest) Test_deve_remover_um_customer_por_id() {
	adapter := NewCustomerAdapter(suite.db)
	customer, err := entities.CreateCustomer(suite.createCustomerCommand)
	suite.Nil(err)

	err = adapter.Save(*customer)
	suite.Nil(err)

	err = adapter.Delete(customer.GetId())
	suite.Nil(err)

	customerFound, err := adapter.FindById(customer.GetId())
	suite.Nil(customerFound)
	suite.Error(err)
}

func Test_CustomerAdapterSuite(t *testing.T) {
	suite.Run(t, new(CustomerAdapterSuiteTest))
}
