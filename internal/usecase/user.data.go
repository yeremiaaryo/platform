package usecase

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (uu *userUC) RegisterUser(ctx context.Context, user entity.UserInfo) error {
	return uu.userSvc.RegisterUser(ctx, user)
}

func (uu *userUC) ValidateLogin(ctx context.Context, user entity.UserInfo) error {
	return uu.userSvc.ValidateLogin(ctx, user)
}

func (uu *userUC) RefreshCookie(ctx context.Context, cookie string) error {
	return uu.userSvc.RefreshCookie(ctx, cookie)
}

func (uu *userUC) GenerateForgotPasswordToken(ctx context.Context, email string) error {
	return uu.userSvc.GenerateForgotPasswordToken(ctx, email)
}
