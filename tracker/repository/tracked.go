package repository

import (
	//"tracker/models"

	"github.com/jmoiron/sqlx"
)

type trackedRepository struct {
	db *sqlx.DB
}

type TrackedRepository interface {
}

func InitTrackedRepository(db *sqlx.DB) TrackedRepository {
	return &trackedRepository{
		db,
	}
}
