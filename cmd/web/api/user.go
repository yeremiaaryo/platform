package api

import (
	"encoding/json"
	"net/http"

	"github.com/yeremiaaryo/go-pkg/response"
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

func (a *API) ValidateUser(r *http.Request) *response.JSONResponse {
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
