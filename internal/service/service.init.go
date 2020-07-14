package service

import (
	"github.com/yeremiaaryo/go-pkg/crypto"
	"github.com/yeremiaaryo/platform/internal/repository"
)

type userSvc struct {
	userRepo    repository.UserRepository
	hashManager HashManager
	cacheRepo   repository.CacheRepository
}

type HashManager interface {
	GenerateHashedPassword(pwd []byte) (string, error)
	ComparedPassword(hashedPassword, pwd []byte) bool
}

func NewUserService(userRepo repository.UserRepository, crypto crypto.HashManager, cacheRepo repository.CacheRepository) *userSvc {
	return &userSvc{
		userRepo:    userRepo,
		hashManager: crypto,
		cacheRepo:   cacheRepo,
	}
}
