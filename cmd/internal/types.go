package internal

import (
	"github.com/yeremiaaryo/platform/internal/auth"
	"github.com/yeremiaaryo/platform/internal/usecase"
)

type Usecase struct {
	User usecase.UserUsecase
	Auth auth.AuthUsecase
	Shop usecase.ShopUsecase
	Chat usecase.ChatUsecase
}
