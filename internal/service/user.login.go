package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yeremiaaryo/go-pkg/router"
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

	uuid := uuid.New()

	c := &http.Cookie{}
	c.Name = "_SID_HobbyLobby_"
	c.Value = uuid.String()
	c.Expires = time.Now().Add(5 * time.Minute)

	w := router.GetResponseWriter(ctx)
	http.SetCookie(w, c)

	return nil
}
