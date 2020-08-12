package usecase

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (su *shopUC) GetShopInfoByUserID(ctx context.Context, userID int64) (*entity.ShopInfo, error) {
	return su.shopSvc.GetShopInfoByUserID(ctx, userID)
}

func (su *shopUC) InsertUpdateShopData(ctx context.Context, data *entity.ShopInfoRequest) error {
	return su.shopSvc.InsertUpdateShopData(ctx, data)
}

func (su *shopUC) UploadImage(ctx context.Context, image []byte) (interface{}, error) {
	return su.shopSvc.UploadImage(ctx, image)
}
