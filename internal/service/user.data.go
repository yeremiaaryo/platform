package service

import (
	"context"
	"errors"
	"log"

	"github.com/yeremiaaryo/platform/internal/entity"
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

func (us *userSvc) RegisterUser(ctx context.Context, user entity.UserInfo) error {
	err := us.userRepo.RegisterUser(ctx, user)
	if err != nil {
		log.Println("Error when register user", err.Error())
		return err
	}
	return nil
}
