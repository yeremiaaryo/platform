package service

import (
	"context"
	"encoding/json"
	"log"

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
	skills := []entity.ShopSkill{}
	_ = json.Unmarshal(data.Skills, &skills)
	return &entity.ShopInfo{
		ID:                data.ID,
		UserID:            data.UserID,
		ShopName:          data.ShopName,
		ShopBanner:        data.ShopBanner.String,
		Description:       data.Description.String,
		CurrentOccupation: data.CurrentOccupation.String,
		Skills:            skills,
		PersonalWebsite:   data.PersonalWebsite.String,
		InstagramURL:      data.InstagramURL.String,
		Category:          data.Category.String,
	}
}

func (ss *shopSvc) InsertUpdateShopData(ctx context.Context, data *entity.ShopInfoRequest) error {
	shop, err := ss.shopRepo.GetShopInfoByUserID(ctx, data.UserID)
	if err != nil {
		return err
	}

	skills, _ := json.Marshal(data.Skills)
	data.SkillsDB = skills

	if shop == nil {
		err = ss.shopRepo.InsertShopData(ctx, data)
		if err != nil {
			log.Println("Error insert shop data:", err.Error())
			return err
		}
		return nil
	}

	err = ss.shopRepo.UpdateShopData(ctx, data)
	if err != nil {
		log.Println("Error update shop data:", err.Error())
		return err
	}
	return nil
}
