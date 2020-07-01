package repository

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

type UserRepository interface {
	FetchUserDataByEmail(ctx context.Context, email string) (*entity.UserInfo, error)
	RegisterUser(ctx context.Context, user entity.UserInfo) error
}
