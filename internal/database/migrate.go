package database

import (
	"github.com/pressly/goose/v3"
)

func RunMigrations() error {
	gormDB := GetSqlDb()
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}

	// Set the dialect
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	// Run migrations
	return goose.Up(sqlDB, "migrations")
}

func RollbackLastMigration() error {
	gormDB := GetSqlDb()
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	return goose.Down(sqlDB, "migrations")
}
