package entity

import "database/sql"

type ShopInfoDB struct {
	ID          int64          `json:"id" db:"id"`
	UserID      int64          `json:"user_id" db:"user_id"`
	ShopName    string         `json:"shop_name" db:"shop_name"`
	ShopAvatar  sql.NullString `json:"shop_avatar" db:"shop_avatar"`
	Description sql.NullString `json:"description" db:"description"`
	Tagline     sql.NullString `json:"tagline" db:"tagline"`
	Category    sql.NullString `json:"category" db:"category"`
}

type ShopInfo struct {
	ID          int64  `json:"id" db:"id"`
	UserID      int64  `json:"user_id" db:"user_id"`
	ShopName    string `json:"shop_name" db:"shop_name"`
	ShopAvatar  string `json:"shop_avatar" db:"shop_avatar"`
	Description string `json:"description" db:"description"`
	Tagline     string `json:"tagline" db:"tagline"`
	Category    string `json:"category" db:"category"`
}
