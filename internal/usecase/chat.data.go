package usecase

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (cu *chatUC) InsertChatOrder(ctx context.Context, data entity.OrderChatRequest, userID int64) error {
	return cu.chatSvc.InsertChatOrder(ctx, data, userID)
}
