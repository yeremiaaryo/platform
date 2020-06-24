package usecase

import "context"

type UserUsecase interface {
	GetUserName(ctx context.Context, userID int64) (string, error)
}
