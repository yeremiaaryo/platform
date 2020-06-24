package service

import "github.com/yeremiaaryo/platform/internal/repository"

type userSvc struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userSvc {
	return &userSvc{
		userRepo: userRepo,
	}
}
