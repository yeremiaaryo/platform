package service

import (
	"context"
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

		_ = json.Unmarshal(v.Tags, tags)
		_ = json.Unmarshal(v.Category, category)
		_ = json.Unmarshal(v.ImageURL, imageURL)
		if v.CatalogID.Valid {
			catalogID = v.CatalogID.Int64
		}
		resp := entity.InspirationList{
			ID:          v.ID,
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
