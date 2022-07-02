package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/go-api-server/internal/entity"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password) values ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetUser(email string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id, password FROM %s WHERE email=$1", userTable)
	err := r.db.Get(&user, query, email)

	return user, err
}
