package repository

import (
	"database/sql"
	"local.packages/domain"
)

type UserRepository interface {
	Insert(DB *sql.DB, userId,name,email string) error
	GetByUserID(DB *sql.DB, userId string) ([]*domain.User, error)
}