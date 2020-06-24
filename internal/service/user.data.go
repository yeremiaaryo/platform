package service

import (
	"context"
	"errors"
)

func (us *userSvc) GetUserName(ctx context.Context, userID int64) (string, error) {
	userInfo, err := us.userRepo.FetchUserDataByUserID(ctx, userID)
	if err != nil {
		return "", err
	}

	if userInfo == nil {
		return "", errors.New("Empty user data")
	}
	return userInfo.Name, nil
}
