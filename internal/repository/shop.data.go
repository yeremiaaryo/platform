package repository

import (
	"context"
	"database/sql"

	"github.com/yeremiaaryo/platform/internal/entity"
)

const getShopInfo = `SELECT id, user_id, shop_name, shop_banner, description, current_occupation, skills, personal_website, instagram_url, category FROM shop WHERE user_id = ?`

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

const insertShopData = `INSERT INTO shop (user_id, shop_name, shop_banner, description, current_occupation, skills, personal_website, instagram_url, category) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`

func (sr *shopRepo) InsertShopData(ctx context.Context, data *entity.ShopInfoRequest) error {
	_, err := sr.db.GetMaster().ExecContext(ctx, insertShopData, data.UserID, data.ShopName, data.ShopBanner, data.Description, data.CurrentOccupation, data.SkillsDB, data.PersonalWebsite, data.InstagramURL, data.Category)
	return err
}

const updateShopData = `UPDATE shop SET shop_name = ?, shop_banner = ?, description = ?, current_occupation = ?, skills = ?, personal_website = ?, instagram_url = ?, category = ? WHERE user_id = ?;
`

func (sr *shopRepo) UpdateShopData(ctx context.Context, data *entity.ShopInfoRequest) error {
	_, err := sr.db.GetMaster().ExecContext(ctx, updateShopData, data.ShopName, data.ShopBanner, data.Description, data.CurrentOccupation, data.SkillsDB, data.PersonalWebsite, data.InstagramURL, data.Category, data.UserID)
	return err
}
