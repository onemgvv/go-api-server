package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/go-api-server/internal/entity"
)

const (
	userTable = "users"
)

type User interface {
	CreateUser(user entity.User) (int, error)
	GetUser(email string) (entity.User, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
