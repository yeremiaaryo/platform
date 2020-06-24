package service

import "context"

type UserService interface {
	GetUserName(ctx context.Context, userID int64) (string, error)
}
