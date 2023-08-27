package main

import (
	"venda-de-ingressos/configs"
)

func main() {
	env := configs.LoadEnvConfig("./cmd/server/.env")
	configs.LoadMigrationUp(*env)
	configs.LoadDataBase(*env)
}
