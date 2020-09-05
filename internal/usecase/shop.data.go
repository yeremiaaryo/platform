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

func (su *shopUC) UploadImage(ctx context.Context, image []byte, folder string) (*entity.UploadImageResponse, error) {
	return su.shopSvc.UploadImage(ctx, image, folder)
}

func (su *shopUC) GetInspirationList(ctx context.Context, userID int64) ([]entity.InspirationList, error) {
	return su.shopSvc.GetInspirationList(ctx, userID)
}

func (su *shopUC) InsertInspiration(ctx context.Context, data entity.InspirationList, userID int64) error {
	return su.shopSvc.InsertInspiration(ctx, data, userID)
}
