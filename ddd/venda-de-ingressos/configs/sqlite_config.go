package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func LoadSqlite(connectionString string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Printf("banco de dados %s conectado com sucesso", db.Name())
	return db
}
