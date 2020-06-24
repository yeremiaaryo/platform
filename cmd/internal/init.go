package internal

import (
	"github.com/yeremiaaryo/platform/internal/repository"
	"github.com/yeremiaaryo/platform/internal/service"
	"github.com/yeremiaaryo/platform/internal/usecase"

	"github.com/yeremiaaryo/go-pkg/database"
)

func GetUsecase(db *database.Store) *Usecase {
	userRepo := repository.NewUserRepo(db)

	userSvc := service.NewUserService(userRepo)

	userUC := usecase.NewUserUsecase(userSvc)

	return &Usecase{
		User: userUC,
	}
}
