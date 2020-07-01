package service

import (
	"github.com/yeremiaaryo/go-pkg/crypto"
	"github.com/yeremiaaryo/platform/internal/repository"
)

type userSvc struct {
	userRepo    repository.UserRepository
	hashManager HashManager
}

type HashManager interface {
	GenerateHashedPassword(pwd []byte) (string, error)
	ComparedPassword(hashedPassword, pwd []byte) bool
}

func NewUserService(userRepo repository.UserRepository, crypto crypto.HashManager) *userSvc {
	return &userSvc{
		userRepo:    userRepo,
		hashManager: crypto,
	}
}
