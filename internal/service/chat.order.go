package service

import (
	"context"
	"log"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (cs *chatSvc) InsertChatOrder(ctx context.Context, data entity.OrderChatRequest, userID int64) error {
	orderChat, err := cs.chatRepo.GetOrderChat(ctx, data.InvoiceNo)
	if err != nil {
		log.Println("[chatSvc][InsertChatOrder] Error on getting order chat", err.Error())
		return err
	}

	if orderChat != nil && orderChat.Status == entity.ChatStatusOpen {
		historyDetails := entity.OrderChatHistory{
			OrderChatID: orderChat.ID,
			Content:     data.Content,
			UserID:      userID,
		}
		err = cs.chatRepo.InsertOrderChatHistory(ctx, historyDetails)
		if err != nil {
			log.Println("[chatSvc][InsertChatOrder] Error insert order chat history, order chat exist", err.Error())
			return err
		}
	} else {
		id, err := cs.chatRepo.InsertOrderChat(ctx, data.InvoiceNo, userID)
		if err != nil {
			log.Println("[chatSvc][InsertChatOrder] Error insert order chat", err.Error())
			return err
		}

		historyDetails := entity.OrderChatHistory{
			OrderChatID: id,
			Content:     data.Content,
			UserID:      userID,
		}
		err = cs.chatRepo.InsertOrderChatHistory(ctx, historyDetails)
		if err != nil {
			log.Println("[chatSvc][InsertChatOrder] Error insert order chat history, order chat not exist", err.Error())
			return err
		}
	}
	return err
}
