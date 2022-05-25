package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

func ConnectDatabase() {
	dsn := "host=localhost user=sharan password=pass dbname=practice port=5432 "
	DB, Err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if Err != nil {
		panic(fmt.Sprintf("Error while connecting DB {{.}}", Err))
	}

	// schema migration
	DB.AutoMigrate(&Code{})

	fmt.Println("working till here", DB)

}

func GetDB() *gorm.DB {
	return DB
}
