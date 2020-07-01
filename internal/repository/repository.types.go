package repository

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

type UserRepository interface {
	FetchUserDataByUserID(ctx context.Context, userID int64) (*entity.UserInfo, error)
	RegisterUser(ctx context.Context, user entity.UserInfo) error
}
