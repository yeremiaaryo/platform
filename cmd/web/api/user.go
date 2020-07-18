package api

import (
	"encoding/json"
	"net/http"

	"github.com/yeremiaaryo/go-pkg/response"
	"github.com/yeremiaaryo/platform/internal/auth"
	"github.com/yeremiaaryo/platform/internal/entity"
)

func (a *API) RegisterUser(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	user := entity.UserInfo{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	err = a.userUC.RegisterUser(ctx, user)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse()
}

func (a *API) ValidateLogin(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	user := entity.UserInfo{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	err = a.userUC.ValidateLogin(ctx, user)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	return response.NewJSONResponse()
}

func (a *API) ValidateCookie(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	cookie, err := auth.GetCookieValue(r, entity.CookieName)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	err = a.userUC.RefreshCookie(ctx, cookie)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrInternalServerError).SetMessage(err.Error())
	}
	return response.NewJSONResponse()
}

func (a *API) GenerateForgotPasswordToken(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	_, email := auth.GetUserDetailFromContext(ctx)
	err := a.userUC.GenerateForgotPasswordToken(ctx, email)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse()
}

func (a *API) ValidateForgotPasswordToken(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	_, email := auth.GetUserDetailFromContext(ctx)
	otpValue := r.Header.Get("OTP-Val")
	correct, err := a.userUC.ValidateForgotPasswordToken(ctx, email, otpValue)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse().SetData(correct)
}

func (a *API) ResetPassword(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	data := entity.ResetPassword{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	_, email := auth.GetUserDetailFromContext(ctx)
	err = a.userUC.ResetPassword(ctx, data, email)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse()
}
