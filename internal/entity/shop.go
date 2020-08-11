package entity

import "database/sql"

type ShopInfoDB struct {
	ID                int64          `json:"id" db:"id"`
	UserID            int64          `json:"user_id" db:"user_id"`
	ShopName          string         `json:"shop_name" db:"shop_name"`
	ShopBanner        sql.NullString `json:"shop_banner" db:"shop_banner"`
	Description       sql.NullString `json:"description" db:"description"`
	CurrentOccupation sql.NullString `json:"current_occupation" db:"current_occupation"`
	Skills            []byte         `json:"skills" db:"skills"`
	PersonalWebsite   sql.NullString `json:"personal_website" db:"personal_website"`
	InstagramURL      sql.NullString `json:"instagram_url" db:"instagram_url"`
	Category          sql.NullString `json:"category" db:"category"`
}

type ShopInfo struct {
	ID                int64       `json:"id" db:"id"`
	UserID            int64       `json:"user_id" db:"user_id"`
	ShopName          string      `json:"shop_name" db:"shop_name"`
	ShopBanner        string      `json:"shop_banner" db:"shop_banner"`
	Description       string      `json:"description" db:"description"`
	CurrentOccupation string      `json:"current_occupation" db:"current_occupation"`
	Skills            []ShopSkill `json:"skills" db:"skills"`
	PersonalWebsite   string      `json:"personal_website" db:"personal_website"`
	InstagramURL      string      `json:"instagram_url" db:"instagram_url"`
	Category          string      `json:"category" db:"category"`
}

type ShopInfoRequest struct {
	UserID            int64       `json:"user_id" db:"user_id"`
	ShopName          string      `json:"shop_name" db:"shop_name"`
	ShopBanner        string      `json:"shop_banner" db:"shop_banner"`
	Description       string      `json:"description" db:"description"`
	CurrentOccupation string      `json:"current_occupation" db:"current_occupation"`
	Skills            []ShopSkill `json:"skills" db:"skills"`
	PersonalWebsite   string      `json:"personal_website" db:"personal_website"`
	InstagramURL      string      `json:"instagram_url" db:"instagram_url"`
	Category          string      `json:"category" db:"category"`
	SkillsDB          []byte
}

type ShopSkill struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}
