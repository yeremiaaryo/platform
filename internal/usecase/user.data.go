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

func (uu *userUC) ValidateForgotPasswordToken(ctx context.Context, email, token string) (bool, error) {
	return uu.userSvc.ValidateForgotPasswordToken(ctx, email, token)
}

func (uu *userUC) ResetPassword(ctx context.Context, data entity.ResetPassword, email string) error {
	return uu.userSvc.ResetPassword(ctx, data, email)
}
