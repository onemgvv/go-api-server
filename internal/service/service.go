package service

import (
	"github.com/onemgvv/go-api-server/internal/entity"
	"github.com/onemgvv/go-api-server/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
}

type Service struct {
	Authorization
	User
}

type Deps struct {
	Repos *repository.Repository
}

func NewService(deps *Deps) *Service {
	authService := NewAuthService(deps.Repos.User)
	return &Service{
		Authorization: authService,
	}
}
