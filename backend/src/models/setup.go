package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	Err error
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=sharan password=pass dbname=practice port=5432 "
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
