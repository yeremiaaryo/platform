package usecase

import (
	"github.com/yeremiaaryo/platform/internal/service"
)

type userUC struct {
	userSvc service.UserService
}

func NewUserUsecase(userSvc service.UserService) *userUC {
	return &userUC{
		userSvc: userSvc,
	}
}

type shopUC struct {
	shopSvc service.ShopService
}

func NewShopUsecase(shopSvc service.ShopService) *shopUC {
	return &shopUC{
		shopSvc: shopSvc,
	}
}

type chatUC struct {
	chatSvc service.ChatService
}

func NewChatUsecase(chatSvc service.ChatService) *chatUC {
	return &chatUC{
		chatSvc: chatSvc,
	}
}
