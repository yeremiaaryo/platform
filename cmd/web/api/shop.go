package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

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

func (a *API) UploadImage(r *http.Request) *response.JSONResponse {
	const fiveMB = 5 << 20
	ctx := r.Context()

	r.Body = http.MaxBytesReader(httptest.NewRecorder(), r.Body, fiveMB)
	if err := r.ParseMultipartForm(fiveMB); err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	img, err := ioutil.ReadAll(io.LimitReader(file, fiveMB))
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	folder := r.FormValue("type")
	result, err := a.shopUC.UploadImage(ctx, img, folder)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse().SetData(result)
}

func (a *API) GetInspirationList(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	userID, _ := auth.GetUserDetailFromContext(ctx)
	inspiration, err := a.shopUC.GetInspirationList(ctx, userID)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse().SetData(inspiration)
}

func (a *API) InsertInspiration(r *http.Request) *response.JSONResponse {
	ctx := r.Context()

	data := entity.InspirationList{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	userID, _ := auth.GetUserDetailFromContext(ctx)
	err = a.shopUC.InsertInspiration(ctx, data, userID)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	return response.NewJSONResponse()
}
