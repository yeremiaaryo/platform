package usecase

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, user entity.UserInfo) error
	ValidateLogin(ctx context.Context, user entity.UserInfo) error
	RefreshCookie(ctx context.Context, cookie string) error
	GenerateForgotPasswordToken(ctx context.Context, email string) error
	ValidateForgotPasswordToken(ctx context.Context, email, token string) (bool, error)
	ResetPassword(ctx context.Context, data entity.ResetPassword, email string) error
	ValidateVerifyToken(ctx context.Context, jwtToken string) error
	IsVerified(ctx context.Context, email string) (bool, error)
	ResendVerificationEmail(ctx context.Context, userID int64, email string) error
}

type ShopUsecase interface {
	GetShopInfoByUserID(ctx context.Context, userID int64) (*entity.ShopInfo, error)
	InsertUpdateShopData(ctx context.Context, data *entity.ShopInfoRequest) error
	UploadImage(ctx context.Context, image []byte, folder string) (*entity.UploadImageResponse, error)
	GetInspirationList(ctx context.Context, userID int64) ([]entity.InspirationList, error)
	InsertInspiration(ctx context.Context, data entity.InspirationList, userID int64) error
}

type ChatUsecase interface {
	InsertChatOrder(ctx context.Context, data entity.OrderChatRequest, userID int64) error
	GetOrderChatHistoryList(ctx context.Context, invoiceNo string) ([]entity.OrderChatHistoryList, error)
}
