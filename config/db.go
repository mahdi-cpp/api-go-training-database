package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var DB *gorm.DB
var err error

func DatabaseInit() {
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=PostgreSQL user=mahdi password=aliali dbname=training port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "api_v1.", // schema name
			SingularTable: false,
		}})

	if err != nil {
		println("Failed to connect database gallery\"")
		os.Exit(1)
	}
}
