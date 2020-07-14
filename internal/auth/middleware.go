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

		c, err := r.Cookie(entity.CookieName)
		if err != nil {
			return response.NewJSONResponse().SetError(response.ErrForbiddenResource).SetMessage(err.Error())
		}

		cookie := c.Value
		userID, err := au.userSvc.ValidateCookies(ctx, cookie)
		if err != nil {
			return response.NewJSONResponse().SetError(response.ErrForbiddenResource).SetMessage(err.Error())
		}

		ctx = context.WithValue(ctx, entity.ContextUserID, userID)
		r = r.WithContext(ctx)
		return h(r)
	}
}
