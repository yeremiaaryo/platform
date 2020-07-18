package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/yeremiaaryo/platform/internal/entity"
	"github.com/yeremiaaryo/platform/internal/utils"
	"gopkg.in/gomail.v2"
)

func (us *userSvc) GenerateForgotPasswordToken(ctx context.Context, email string) error {
	user, err := us.userRepo.FetchUserDataByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("Invalid user email")
	}

	otpToken, err := utils.GenerateOTP()
	if err != nil {
		return err
	}

	redisKey := fmt.Sprintf(entity.RedisKeyForgotPasswordToken, email)
	err = us.cacheRepo.Set(redisKey, otpToken, entity.OTPExpiredInSeconds)
	if err != nil {
		return err
	}

	message := fmt.Sprintf(`Hello, %s! <b>You have requested to reset your password</b><br/>Your token is %s`, user.Name, otpToken)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", entity.ConfigEmail)
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", "Forgot Password Request")
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		entity.ConfigSMTPHost,
		entity.ConfigSMTPPort,
		entity.ConfigEmail,
		entity.ConfigPassword,
	)

	go func(mailer *gomail.Message) {
		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Println("Error sending email", err.Error())
			return
		}
		log.Println("Success sending email")
	}(mailer)

	return nil
}

func (us *userSvc) ValidateForgotPasswordToken(ctx context.Context, email, token string) (bool, error) {
	user, err := us.userRepo.FetchUserDataByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, errors.New("Invalid user email")
	}

	redisKey := fmt.Sprintf(entity.RedisKeyForgotPasswordToken, email)
	redisToken, err := us.cacheRepo.Get(redisKey)
	if err != nil {
		return false, err
	}

	if redisToken != token {
		return false, nil
	}

	return true, nil
}

func (us *userSvc) ResetPassword(ctx context.Context, data entity.ResetPassword, email string) error {
	user, err := us.userRepo.FetchUserDataByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("Invalid user email")
	}

	isPasswordCorrect := us.hashManager.ComparedPassword([]byte(user.Password), []byte(data.OldPassword))
	if !isPasswordCorrect {
		return errors.New("Invalid Old Password")
	}

	hashedPwd, err := us.hashManager.GenerateHashedPassword([]byte(data.NewPassword))
	if err != nil {
		return err
	}
	data.NewPassword = hashedPwd

	return us.userRepo.ResetPassword(ctx, data, email)
}
