package service

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

type UserService interface {
	RegisterUser(ctx context.Context, user entity.UserInfo) error
	ValidateLogin(ctx context.Context, data entity.UserInfo) error
}
