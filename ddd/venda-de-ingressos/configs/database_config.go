package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func LoadDataBase(env EnvConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(env.DBDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Printf("banco de dados %s conectado com sucesso", db.Name())
	return db
}
