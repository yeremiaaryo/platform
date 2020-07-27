package service

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (ss *shopSvc) GetShopInfoByUserID(ctx context.Context, userID int64) (*entity.ShopInfo, error) {
	shop, err := ss.shopRepo.GetShopInfoByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if shop == nil {
		return nil, nil
	}
	return buildShopResponse(shop), nil
}

func buildShopResponse(data *entity.ShopInfoDB) *entity.ShopInfo {
	return &entity.ShopInfo{
		ID:          data.ID,
		UserID:      data.UserID,
		ShopName:    data.ShopName,
		ShopAvatar:  data.ShopAvatar.String,
		Description: data.Description.String,
		Tagline:     data.Tagline.String,
		Category:    data.Category.String,
	}
}
