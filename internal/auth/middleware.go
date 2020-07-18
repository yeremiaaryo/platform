package auth

import (
	"context"
	"net/http"

	"github.com/yeremiaaryo/go-pkg/response"
	"github.com/yeremiaaryo/go-pkg/router"
	"github.com/yeremiaaryo/platform/internal/entity"
)

func (au *authUC) Authorize(h router.Handle) router.Handle {
	return func(r *http.Request) *response.JSONResponse {
		ctx := r.Context()

		cookie, err := GetCookieValue(r, entity.CookieName)
		if err != nil {
			return response.NewJSONResponse().SetError(response.ErrForbiddenResource).SetMessage(err.Error())
		}

		userID, email, err := au.userSvc.ValidateCookies(ctx, cookie)
		if err != nil {
			return response.NewJSONResponse().SetError(response.ErrForbiddenResource).SetMessage(err.Error())
		}

		ctx = context.WithValue(ctx, entity.ContextUserID, userID)
		ctx = context.WithValue(ctx, entity.ContextEmail, email)
		r = r.WithContext(ctx)
		return h(r)
	}
}

func GetUserDetailFromContext(ctx context.Context) (int64, string) {
	uID := ctx.Value(entity.ContextUserID)
	if uID == nil {
		return 0, ""
	}
	userID := uID.(int64)

	e := ctx.Value(entity.ContextEmail)
	if e == nil {
		return 0, ""
	}
	email := e.(string)

	return userID, email
}

func GetCookieValue(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(entity.CookieName)
	if err != nil {
		return "", err
	}
	return c.Value, nil
}
