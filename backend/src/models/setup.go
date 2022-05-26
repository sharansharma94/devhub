package models

import (
	"devhub/src/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	Err error
)

func ConnectDatabase() *gorm.DB {
	conf := config.GetDBConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", conf.POSTGRES_HOST, conf.POSTGRES_USER, conf.POSTGRES_PASSWORD, conf.POSTGRES_DB, conf.POSTGRES_PORT)
	fmt.Println(dsn)
	Db, Err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if Err != nil {
		panic(fmt.Sprintf("Error while connecting DB {{.}}", Err))
	}

	// schema migration
	Db.AutoMigrate(&Code{})
	return Db
}

func GetDB() *gorm.DB {
	if Db == nil {
		Db = ConnectDatabase()
	}
	return Db
}
