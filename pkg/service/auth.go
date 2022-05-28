package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/onemgvv/go-api-server/pkg/entity"
	"github.com/onemgvv/go-api-server/pkg/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	signingKey = "wq#diu238174y9jbf203ce#feif"
	tokenTTL   = 12 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		return "", err
	}

	_, checkError := checkPasswordHash(password, user.Password)
	if checkError != nil {
		return "", checkError
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func generatePasswordHash(value string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 7)
	if err != nil {
		logrus.Error("password hashing error: %s", err.Error())
	}

	return string(bytes)
}

func checkPasswordHash(value string, hash string) (bool, error) {
	plainPassword := []byte(value)
	err := bcrypt.CompareHashAndPassword([]byte(hash), plainPassword)
	if err != nil {
		return false, err
	}

	return true, nil
}
