package usecase

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, user entity.UserInfo) error
	ValidateLogin(ctx context.Context, user entity.UserInfo) error
	RefreshCookie(ctx context.Context, cookie string) error
	GenerateForgotPasswordToken(ctx context.Context, email string) error
}
