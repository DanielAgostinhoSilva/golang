package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"venda-de-ingressos/configs"
)

func main() {
	env := configs.LoadEnvConfig("./cmd/server/.env")
	configs.LoadMigration(*env)
	db, err := gorm.Open(mysql.Open(env.DBDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Printf("banco de dados %s conectado com sucesso", db.Name())
}
