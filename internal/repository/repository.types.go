package repository

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

type UserRepository interface {
	FetchUserDataByEmail(ctx context.Context, email string) (*entity.UserInfo, error)
	RegisterUser(ctx context.Context, user entity.UserInfo) error
}

type CacheRepository interface {
	Get(key string) (string, error)
	Set(key, value string, expired int) error
	Del(key string) error
}
