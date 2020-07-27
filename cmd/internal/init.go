package internal

import (
	"github.com/yeremiaaryo/platform/internal/auth"
	"github.com/yeremiaaryo/platform/internal/repository"
	"github.com/yeremiaaryo/platform/internal/service"
	"github.com/yeremiaaryo/platform/internal/usecase"

	"github.com/yeremiaaryo/go-pkg/crypto"
	"github.com/yeremiaaryo/go-pkg/database"
)

func GetUsecase(db *database.Store, cache repository.RedisConn) *Usecase {
	hashManager := crypto.NewHashManager()

	cacheRepo := repository.NewCacheRepo(cache)
	userRepo := repository.NewUserRepo(db)
	shopRepo := repository.NewShopRepo(db)

	userSvc := service.NewUserService(userRepo, hashManager, cacheRepo)
	shopSvc := service.NewShopService(userRepo, shopRepo)

	userUC := usecase.NewUserUsecase(userSvc)
	authUC := auth.NewAuthUsecase(userSvc)
	shopUC := usecase.NewShopUsecase(shopSvc)

	return &Usecase{
		User: userUC,
		Auth: authUC,
		Shop: shopUC,
	}
}
