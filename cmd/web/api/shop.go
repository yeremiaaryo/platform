package api

import (
	"encoding/json"
	"net/http"

	"github.com/yeremiaaryo/go-pkg/response"
	"github.com/yeremiaaryo/platform/internal/auth"
	"github.com/yeremiaaryo/platform/internal/entity"
)

func (a *API) GetShopInfo(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	userID, _ := auth.GetUserDetailFromContext(ctx)
	shop, err := a.shopUC.GetShopInfoByUserID(ctx, userID)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse().SetData(shop)
}

func (a *API) InsertUpdateShopData(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	data := entity.ShopInfoRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	userID, _ := auth.GetUserDetailFromContext(ctx)
	data.UserID = userID
	err = a.shopUC.InsertUpdateShopData(ctx, &data)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse()
}
