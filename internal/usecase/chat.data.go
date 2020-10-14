package usecase

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (cu *chatUC) InsertChatOrder(ctx context.Context, data entity.OrderChatRequest, userID int64) error {
	return cu.chatSvc.InsertChatOrder(ctx, data, userID)
}

func (cu *chatUC) GetOrderChatHistoryList(ctx context.Context, invoiceNo string) ([]entity.OrderChatHistoryList, error) {
	return cu.chatSvc.GetOrderChatHistoryList(ctx, invoiceNo)
}
