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
	cloudinaryRepo := repository.NewCloudinaryRepo()
	inspirationRepo := repository.NewInspirationRepo(db)
	chatRepo := repository.NewChatRepo(db)

	userSvc := service.NewUserService(userRepo, hashManager, cacheRepo)
	shopSvc := service.NewShopService(userRepo, shopRepo, cloudinaryRepo, inspirationRepo)
	chatSvc := service.NewChatService(chatRepo)

	userUC := usecase.NewUserUsecase(userSvc)
	authUC := auth.NewAuthUsecase(userSvc)
	shopUC := usecase.NewShopUsecase(shopSvc)
	chatUC := usecase.NewChatUsecase(chatSvc)

	return &Usecase{
		User: userUC,
		Auth: authUC,
		Shop: shopUC,
		Chat: chatUC,
	}
}
