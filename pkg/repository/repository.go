package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/go-api-server/pkg/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(email string) (entity.User, error)
}

type Users interface {
}

type User interface {
}

type Repository struct {
	Authorization
	Users
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
