package repository

import (
	"context"
	"database/sql"

	"github.com/yeremiaaryo/platform/internal/entity"
)

const getInspirationList = `SELECT id, shop_id, title, tags, category, description, image_url, catalog_id FROM inspiration WHERE shop_id = ? AND status = ?`

func (ir *inspirationRepo) GetInspirationListByShopID(ctx context.Context, shopID int64) ([]entity.InspirationListDB, error) {
	resp := []entity.InspirationListDB{}
	err := ir.db.GetSlave().SelectContext(ctx, &resp, getInspirationList, shopID, entity.InspirationActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return resp, err
}

const insertInspiration = `INSERT INTO inspiration (shop_id, title, tags, category, description, image_url, catalog_id) VALUES (?, ?, ?, ?, ?, ?, ?)`

func (ir *inspirationRepo) InsertInspiration(ctx context.Context, data entity.InspirationListDB) error {
	_, err := ir.db.GetSlave().ExecContext(ctx, insertInspiration, data.ShopID, data.Title, data.Tags, data.Category, data.Description, data.ImageURL, data.CatalogID)
	return err
}
