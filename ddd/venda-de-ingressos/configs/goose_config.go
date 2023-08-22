package configs

import (
	"github.com/pressly/goose/v3"
	"log"
)

func LoadMigration(env envConfig) {
	connectionString := env.DBDsn
	dbDriver := env.DBDriver
	migrationsDir := "./internal/infra/db/migration"

	gooseDB, err := goose.OpenDBWithDriver(dbDriver, connectionString)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
		panic(err)
	}

	err = goose.SetDialect(dbDriver)
	if err != nil {
		log.Fatal("Erro ao configurar o dialect:", err)
		panic(err)
	}

	err = goose.Run("up", gooseDB, migrationsDir)
	if err != nil {
		log.Fatal("Erro ao executar as migrações:", err)
		panic(err)
	}
}
