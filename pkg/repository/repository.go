package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
