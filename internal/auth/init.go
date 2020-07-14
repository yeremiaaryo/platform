package auth

import "github.com/yeremiaaryo/platform/internal/service"

type authUC struct {
	userSvc service.UserService
}

func NewAuthUsecase(userSvc service.UserService) *authUC {
	return &authUC{
		userSvc: userSvc,
	}
}
