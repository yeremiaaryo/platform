package usecase

import (
	"github.com/yeremiaaryo/platform/internal/service"
)

type userUC struct {
	userSvc service.UserService
}

func NewUserUsecase(userSvc service.UserService) *userUC {
	return &userUC{
		userSvc: userSvc,
	}
}
