package service

import "github.com/onemgvv/go-api-server/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
