package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SQLDB *gorm.DB

func ConnectToSqlDb() {
	var err error
	dsn := os.Getenv("POSTGRE_CONNECTION_URL")
	SQLDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Postgres")
	}
}

func GetSqlDb() *gorm.DB {
	return SQLDB
}
