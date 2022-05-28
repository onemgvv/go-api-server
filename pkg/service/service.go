package service

import (
	"github.com/onemgvv/go-api-server/pkg/entity"
	"github.com/onemgvv/go-api-server/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(email, password string) (string, error)
}

type Users interface {
}

type User interface {
}

type Service struct {
	Authorization
	Users
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
