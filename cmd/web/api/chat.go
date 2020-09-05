package api

import (
	"encoding/json"
	"net/http"

	"github.com/yeremiaaryo/go-pkg/response"
	"github.com/yeremiaaryo/platform/internal/auth"
	"github.com/yeremiaaryo/platform/internal/entity"
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

func (a *API) GetChatOrderList(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	invoiceNo := r.URL.Query().Get("invoice_no")
	list, err := a.chatUC.GetOrderChatHistoryList(ctx, invoiceNo)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse().SetData(list)
}
