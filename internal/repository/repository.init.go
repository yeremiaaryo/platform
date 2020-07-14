package repository

import "github.com/yeremiaaryo/go-pkg/database"

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
