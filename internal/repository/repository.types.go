package repository

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

type UserRepository interface {
	FetchUserDataByEmail(ctx context.Context, email string) (*entity.UserInfo, error)
	RegisterUser(ctx context.Context, user entity.UserInfo) (int64, error)
	ResetPassword(ctx context.Context, data entity.ResetPassword, email string) error
	UpdateVerifiedUser(ctx context.Context, userID int64) error
}

type CacheRepository interface {
	Get(key string) (string, error)
	Set(key, value string, expired int) error
	Del(key string) error
	TTL(key string) (int64, error)
	GetInt64(key string) (int64, error)
}

type ShopRepository interface {
	GetShopInfoByUserID(ctx context.Context, userID int64) (*entity.ShopInfoDB, error)
	InsertShopData(ctx context.Context, data *entity.ShopInfoRequest) error
	UpdateShopData(ctx context.Context, data *entity.ShopInfoRequest) error
}

type CloudinaryRepository interface {
	UploadImage(ctx context.Context, image, folder string) (*entity.UploadImageResponse, error)
}

type InspirationRepository interface {
	GetInspirationListByShopID(ctx context.Context, shopID int64) ([]entity.InspirationListDB, error)
	InsertInspiration(ctx context.Context, data entity.InspirationListDB) error
}

type ChatRepository interface {
	GetOrderChat(ctx context.Context, invoiceNo string) (*entity.OrderChat, error)
	InsertOrderChat(ctx context.Context, invoiceNo string, userID int64) (int64, error)
	InsertOrderChatHistory(ctx context.Context, data entity.OrderChatHistory) error
	GetOrderChatList(ctx context.Context, orderChatID int64) ([]entity.OrderChatHistoryList, error)
}
