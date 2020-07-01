package service

import (
	"context"
	"errors"
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

	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	const CONFIG_EMAIL = "cobayeremia@gmail.com"
	const CONFIG_PASSWORD = "$Tokopedia789"
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_EMAIL)
	mailer.SetHeader("To", "yeremia.aryo@gmail.com")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_EMAIL,
		CONFIG_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
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
