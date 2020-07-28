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
	ValidateForgotPasswordToken(ctx context.Context, email, token string) (bool, error)
	ResetPassword(ctx context.Context, data entity.ResetPassword, email string) error
	ValidateVerifyToken(ctx context.Context, jwtToken string) error
}

type ShopUsecase interface {
	GetShopInfoByUserID(ctx context.Context, userID int64) (*entity.ShopInfo, error)
}
