package database

import (
	"github.com/fadedead/opengamba_backend/internal/user"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

type PostgresRegistery struct {
	User user.Repository
}

func NewPostgresRegistry(db *gorm.DB) *PostgresRegistery {
	return &PostgresRegistery{
		User: user.NewRepository(db),
	}
}
