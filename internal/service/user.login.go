package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	c.Name = entity.CookieName
	c.Value = uuid.String()
	c.Expires = time.Now().AddDate(0, 0, entity.CookieExpireInDays)

	w := router.GetResponseWriter(ctx)
	http.SetCookie(w, c)

	loginKey := fmt.Sprintf(entity.RedisKeyLogin, c.Value)
	loginValue := fmt.Sprintf("%v~%v", strconv.FormatInt(user.ID, 10), data.Email)
	err = us.cacheRepo.Set(loginKey, loginValue, entity.LoginExpireInSeconds)
	if err != nil {
		return errors.New("Error setting KV to Redis")
	}

	return nil
}

func (us *userSvc) ValidateCookies(ctx context.Context, cookie string) (int64, error) {
	loginKey := fmt.Sprintf(entity.RedisKeyLogin, cookie)
	details, err := us.cacheRepo.Get(loginKey)
	if err != nil {
		return 0, errors.New("Error getting data from Redis")
	}

	detailList := strings.Split(details, "~")
	if len(detailList) < 2 {
		return 0, errors.New("Invalid value, please login again")
	}

	userID, err := strconv.ParseInt(detailList[0], 10, 64)
	if err != nil {
		return 0, errors.New("Invalid User ID from Redis")
	}

	return userID, nil
}
