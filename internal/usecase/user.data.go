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