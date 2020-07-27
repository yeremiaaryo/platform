package api

import (
	"net/http"

	"github.com/yeremiaaryo/go-pkg/response"
	"github.com/yeremiaaryo/platform/internal/auth"
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
