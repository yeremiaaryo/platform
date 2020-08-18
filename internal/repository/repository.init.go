package repository

import (
	"net/http"
	"time"

	"github.com/yeremiaaryo/go-pkg/database"
)

type userRepo struct {
	db *database.Store
}

func NewUserRepo(db *database.Store) *userRepo {
	return &userRepo{
		db: db,
	}
}

type cacheRepo struct {
	cache RedisConn
}

func NewCacheRepo(cache RedisConn) *cacheRepo {
	return &cacheRepo{
		cache: cache,
	}
}

type shopRepo struct {
	db *database.Store
}

func NewShopRepo(db *database.Store) *shopRepo {
	return &shopRepo{
		db: db,
	}
}

type cloudinaryRepo struct {
	client *http.Client
}

func NewCloudinaryRepo() *cloudinaryRepo {
	return &cloudinaryRepo{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type inspirationRepo struct {
	db *database.Store
}

func NewInspirationRepo(db *database.Store) *inspirationRepo {
	return &inspirationRepo{
		db: db,
	}
}
