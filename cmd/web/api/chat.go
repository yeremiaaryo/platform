package api

import (
	"encoding/json"
	"net/http"

	"github.com/yeremiaaryo/platform/internal/auth"
	"github.com/yeremiaaryo/platform/internal/entity"
	"github.com/yeremiaaryo/platform/vendor/github.com/yeremiaaryo/go-pkg/response"
)

func (a *API) InsertOrderData(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	data := entity.OrderChatRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	userID, _ := auth.GetUserDetailFromContext(ctx)
	err = a.chatUC.InsertChatOrder(ctx, data, userID)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse()
}