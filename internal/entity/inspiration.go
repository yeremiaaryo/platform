package entity

import "database/sql"

type InspirationListDB struct {
	ID          int64         `db:"id"`
	Title       string        `db:"title"`
	Tags        []byte        `db:"tags"`
	Category    []byte        `db:"category"`
	Description string        `db:"description"`
	ImageURL    []byte        `db:"image_url"`
	CatalogID   sql.NullInt64 `db:"catalog_id"`
}

type InspirationList struct {
	ID          int64    `db:"id"`
	Title       string   `db:"title"`
	Tags        []string `db:"tags"`
	Category    []string `db:"category"`
	Description string   `db:"description"`
	ImageURL    []string `db:"image_url"`
	CatalogID   int64    `db:"catalog_id"`
}

type InspirationStatus int

const (
	InspirationInactive InspirationStatus = iota
	InspirationActive
)
