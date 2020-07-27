package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/yeremiaaryo/platform/internal/entity"
	"gopkg.in/gomail.v2"
)

func (us *userSvc) RegisterUser(ctx context.Context, user entity.UserInfo) error {
	err := validateUserRegistration(user)
	if err != nil {
		return err
	}

	userInfo, err := us.userRepo.FetchUserDataByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if userInfo != nil {
		log.Println("User already exist")
		return errors.New("Email already registered")
	}

	hashedPwd, err := us.hashManager.GenerateHashedPassword([]byte(user.Password))
	if err != nil {
		return err
	}

	user.Password = hashedPwd
	err = us.userRepo.RegisterUser(ctx, user)
	if err != nil {
		log.Println("Error when register user", err.Error())
		return err
	}

	message := fmt.Sprintf(`Hello, %s! <b>Welcome to HobbyLobby</b>`, user.Name)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", entity.ConfigEmail)
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", "Welcome to HobbyLobby")
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
		log.Println(("Email is sent"))
	}(mailer)
	return nil
}

func validateUserRegistration(inp entity.UserInfo) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(inp.Email) {
		return errors.New("Email not valid")
	}
	if len(inp.Password) < 8 {
		return errors.New("Password must contain at least 8 characters")
	}
	return nil
}
