package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/yeremiaaryo/go-pkg/response"
)

func (a *API) GetUserName(r *http.Request) *response.JSONResponse {
	ctx := r.Context()
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest)
	}

	name, err := a.userUC.GetUserName(ctx, userID)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrInternalServerError).SetMessage(err.Error())
	}
	return response.NewJSONResponse().SetData(fmt.Sprintf("Hello, %s", name))
}
