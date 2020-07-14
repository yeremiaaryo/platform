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

	userSvc := service.NewUserService(userRepo, hashManager, cacheRepo)

	userUC := usecase.NewUserUsecase(userSvc)
	authUC := auth.NewAuthUsecase(userSvc)

	return &Usecase{
		User: userUC,
		Auth: authUC,
	}
}
