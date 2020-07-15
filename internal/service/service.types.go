package service

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

type UserService interface {
	RegisterUser(ctx context.Context, user entity.UserInfo) error
	ValidateLogin(ctx context.Context, data entity.UserInfo) error
	ValidateCookies(ctx context.Context, cookie string) (int64, string, error)
	RefreshCookie(ctx context.Context, cookie string) error
	GenerateForgotPasswordToken(ctx context.Context, email string) error
}
