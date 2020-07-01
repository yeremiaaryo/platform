package usecase

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, user entity.UserInfo) error
}
