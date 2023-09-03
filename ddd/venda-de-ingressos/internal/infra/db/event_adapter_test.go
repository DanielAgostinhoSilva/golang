package db

import (
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"time"
	"venda-de-ingressos/configs"
	"venda-de-ingressos/internal/domain/entities"
)

type EventAdapterSuitTest struct {
	suite.Suite
	db                   *gorm.DB
	env                  *configs.EnvConfig
	createPartnerCommand entities.CreatePartnerCommand
	createEventCommand   entities.CreateEventCommand
}

func (suite *EventAdapterSuitTest) SetupSuite() {
	suite.env = configs.LoadEnvConfig("./../../../cmd/server/test.env")
	suite.db = configs.LoadDataBase(*suite.env)
}

func (suite *EventAdapterSuitTest) SetupTest() {
	configs.LoadMigrationUp(*suite.env)
	suite.createPartnerCommand = entities.CreatePartnerCommand{
		Name: "Partner Name",
	}

	suite.createEventCommand = entities.CreateEventCommand{
		Name:        "Event Name Test",
		Description: "Event Description Test",
		Date:        time.Date(2023, time.August, 18, 0, 0, 0, 0, time.UTC),
	}
}

func (suite *EventAdapterSuitTest) TearDownTest() {
	configs.LoadMigrationDown(*suite.env)
}

func (suite *EventAdapterSuitTest) Test_ao_adicionar_uma_section_deve_persisitir_um_evento() {
	partner, err := entities.CreatePartner(suite.createPartnerCommand)
	suite.Nil(err)

	partnerAdapter := NewPartnerAdapter(suite.db)
	err = partnerAdapter.Save(*partner)
	suite.Nil(err)

	createEventCommand := entities.CreateEventCommand{
		Name:        "Event Name Test",
		Description: "Event Description Test",
		Date:        time.Date(2023, time.August, 18, 0, 0, 0, 0, time.UTC),
		PartnerId:   partner.GetId(),
	}

	event, err := entities.CreateEvent(createEventCommand)
	suite.Nil(err)

	err = event.AddSection(entities.AddSectionCommand{
		Name:        "Section A",
		Description: "Description A",
		TotalSpot:   10,
		Price:       10.00,
	})
	suite.Nil(err)
	err = event.AddSection(entities.AddSectionCommand{
		Name:        "Section B",
		Description: "Description B",
		TotalSpot:   5,
		Price:       15.00,
	})
	suite.Nil(err)

	adapter := NewEventAdapter(suite.db)
	err = adapter.Save(*event)
	suite.Nil(err)

	eventFound, err := adapter.FindById(event.GetId())
	suite.Nil(err)
	suite.Equal(eventFound.GetId(), event.GetId())
	suite.Len(eventFound.GetSections(), 2)
}

func (suite *EventAdapterSuitTest) Test_deve_buscar_todos_os_registro_no_banco_de_dados() {
	partner, err := entities.CreatePartner(suite.createPartnerCommand)
	suite.Nil(err)

	partnerAdapter := NewPartnerAdapter(suite.db)
	err = partnerAdapter.Save(*partner)
	suite.Nil(err)

	createEventCommand := entities.CreateEventCommand{
		Name:        "Event Name Test",
		Description: "Event Description Test",
		Date:        time.Date(2023, time.August, 18, 0, 0, 0, 0, time.UTC),
		PartnerId:   partner.GetId(),
	}

	event, err := entities.CreateEvent(createEventCommand)
	suite.Nil(err)

	adapter := NewEventAdapter(suite.db)
	err = adapter.Save(*event)
	suite.Nil(err)

	events, err := adapter.FindAll()
	suite.Nil(err)
	suite.Len(events, 1)
}

func (suite *EventAdapterSuitTest) Test_deve_remover_um_evento_por_id() {
	partner, err := entities.CreatePartner(suite.createPartnerCommand)
	suite.Nil(err)

	partnerAdapter := NewPartnerAdapter(suite.db)
	err = partnerAdapter.Save(*partner)
	suite.Nil(err)

	createEventCommand := entities.CreateEventCommand{
		Name:        "Event Name Test",
		Description: "Event Description Test",
		Date:        time.Date(2023, time.August, 18, 0, 0, 0, 0, time.UTC),
		PartnerId:   partner.GetId(),
	}

	event, err := entities.CreateEvent(createEventCommand)
	suite.Nil(err)

	adapter := NewEventAdapter(suite.db)
	err = adapter.Save(*event)
	suite.Nil(err)

	err = adapter.Delete(event.GetId())
	suite.Nil(err)

	eventFound, err := adapter.FindById(event.GetId())
	suite.Nil(eventFound)
	suite.Error(err)
}

func Test_EventAdapterSuit(t *testing.T) {
	suite.Run(t, new(EventAdapterSuitTest))
}
