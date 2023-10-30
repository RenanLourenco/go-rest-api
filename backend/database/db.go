package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var (
	DB *gorm.DB
	err error
)

func Connection(){
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(err.Error())
	}
}