package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (ss *shopSvc) GetInspirationList(ctx context.Context, userID int64) ([]entity.InspirationList, error) {
	shop, err := ss.shopRepo.GetShopInfoByUserID(ctx, userID)
	if err != nil {
		log.Println("[GetInspirationList] Error when get shop info", err.Error())
		return nil, err
	}

	if shop == nil {
		return nil, errors.New("Shop hasn't been created")
	}

	listDB, err := ss.inspirationRepo.GetInspirationListByShopID(ctx, shop.ID)
	if err != nil {
		log.Println("[GetInspirationList] Error when get inspiration list", err.Error())
		return nil, err
	}
	return buildInspirationList(listDB), nil
}

func buildInspirationList(data []entity.InspirationListDB) []entity.InspirationList {
	res := []entity.InspirationList{}

	for _, v := range data {
		tags := []string{}
		category := []string{}
		imageURL := []string{}
		catalogID := int64(0)

		_ = json.Unmarshal(v.Tags, &tags)
		_ = json.Unmarshal(v.Category, &category)
		_ = json.Unmarshal(v.ImageURL, &imageURL)
		if v.CatalogID.Valid {
			catalogID = v.CatalogID.Int64
		}
		resp := entity.InspirationList{
			ID:          v.ID,
			ShopID:      v.ShopID,
			Title:       v.Title,
			Tags:        tags,
			Category:    category,
			Description: v.Description,
			ImageURL:    imageURL,
			CatalogID:   catalogID,
		}
		res = append(res, resp)
	}
	return res
}

func (ss *shopSvc) InsertInspiration(ctx context.Context, data entity.InspirationList, userID int64) error {
	shop, err := ss.shopRepo.GetShopInfoByUserID(ctx, userID)
	if err != nil {
		log.Println("[InsertInspiration] Error when get shop info", err.Error())
		return err
	}

	if shop == nil {
		return errors.New("Shop hasn't been created")
	}

	err = ss.inspirationRepo.InsertInspiration(ctx, buildInsertInspirationData(data, shop.ID))
	if err != nil {
		log.Println("[InsertInspiration] error when insert inspiration", err.Error())
		return err
	}
	return nil
}

func buildInsertInspirationData(data entity.InspirationList, shopID int64) entity.InspirationListDB {
	catalogID := sql.NullInt64{}
	if data.CatalogID > 0 {
		catalogID.Int64 = data.CatalogID
		catalogID.Valid = true
	}

	tagsJSON, _ := json.Marshal(data.Tags)
	categoryJSON, _ := json.Marshal(data.Category)
	imageURLJSON, _ := json.Marshal(data.ImageURL)
	resp := entity.InspirationListDB{
		ShopID:      shopID,
		Title:       data.Title,
		Tags:        tagsJSON,
		Category:    categoryJSON,
		Description: data.Description,
		ImageURL:    imageURLJSON,
		CatalogID:   catalogID,
	}
	return resp
}
