package service

import (
	"context"
	"errors"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (us *userSvc) ValidateLogin(ctx context.Context, data entity.UserInfo) error {
	user, err := us.userRepo.FetchUserDataByEmail(ctx, data.Email)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("Email not registered")
	}

	isPasswordCorrect := us.hashManager.ComparedPassword([]byte(user.Password), []byte(data.Password))
	if !isPasswordCorrect {
		return errors.New("Invalid Password")
	}

	return nil
}
