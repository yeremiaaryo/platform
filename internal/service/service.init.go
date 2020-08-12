package service

import (
	"github.com/yeremiaaryo/go-pkg/crypto"
	"github.com/yeremiaaryo/platform/internal/repository"
)

type HashManager interface {
	GenerateHashedPassword(pwd []byte) (string, error)
	ComparedPassword(hashedPassword, pwd []byte) bool
}

type userSvc struct {
	userRepo    repository.UserRepository
	hashManager HashManager
	cacheRepo   repository.CacheRepository
}

func NewUserService(userRepo repository.UserRepository, crypto crypto.HashManager, cacheRepo repository.CacheRepository) *userSvc {
	return &userSvc{
		userRepo:    userRepo,
		hashManager: crypto,
		cacheRepo:   cacheRepo,
	}
}

type shopSvc struct {
	userRepo       repository.UserRepository
	shopRepo       repository.ShopRepository
	cloudinaryRepo repository.CloudinaryRepository
}

func NewShopService(userRepo repository.UserRepository, shopRepo repository.ShopRepository, cloudinaryRepo repository.CloudinaryRepository) *shopSvc {
	return &shopSvc{
		userRepo:       userRepo,
		shopRepo:       shopRepo,
		cloudinaryRepo: cloudinaryRepo,
	}
}
