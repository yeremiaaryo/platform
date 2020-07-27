package repository

import (
	"context"
	"database/sql"

	"github.com/yeremiaaryo/platform/internal/entity"
)

const getShopInfo = `SELECT id, user_id, shop_name, shop_avatar, description, tagline, category FROM shop WHERE user_id = ?`

func (sr *shopRepo) GetShopInfoByUserID(ctx context.Context, userID int64) (*entity.ShopInfoDB, error) {
	resp := new(entity.ShopInfoDB)
	err := sr.db.GetSlave().GetContext(ctx, resp, getShopInfo, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return resp, err
}
