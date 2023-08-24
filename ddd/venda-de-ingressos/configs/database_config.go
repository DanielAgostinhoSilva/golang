package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func LoadDataBase(connectionString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Printf("banco de dados %s conectado com sucesso", db.Name())
	return db
}
