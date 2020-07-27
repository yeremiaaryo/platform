package usecase

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (su *shopUC) GetShopInfoByUserID(ctx context.Context, userID int64) (*entity.ShopInfo, error) {
	return su.shopSvc.GetShopInfoByUserID(ctx, userID)
}
