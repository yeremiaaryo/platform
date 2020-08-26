package entity

import "database/sql"

type InspirationListDB struct {
	ID          int64         `db:"id"`
	ShopID      int64         `db:"shop_id"`
	Title       string        `db:"title"`
	Tags        []byte        `db:"tags"`
	Category    []byte        `db:"category"`
	Description string        `db:"description"`
	ImageURL    []byte        `db:"image_url"`
	CatalogID   sql.NullInt64 `db:"catalog_id"`
}

type InspirationList struct {
	ID          int64    `json:"id"`
	ShopID      int64    `json:"shop_id"`
	Title       string   `json:"title"`
	Tags        []string `json:"tags"`
	Category    []string `json:"category"`
	Description string   `json:"description"`
	ImageURL    []string `json:"image_url"`
	CatalogID   int64    `json:"catalog_id"`
}

type InspirationStatus int

const (
	InspirationInactive InspirationStatus = iota
	InspirationActive
)
