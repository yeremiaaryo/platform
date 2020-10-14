package repository

import (
	"context"
	"database/sql"

	"github.com/yeremiaaryo/platform/internal/entity"
)

const getOrderChat = `SELECT id, invoice_no, status FROM order_chat WHERE invoice_no = ? ORDER BY id DESC`

func (cr *chatRepo) GetOrderChat(ctx context.Context, invoiceNo string) (*entity.OrderChat, error) {
	resp := new(entity.OrderChat)
	err := cr.db.GetSlave().GetContext(ctx, resp, getOrderChat, invoiceNo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}

const insertOrderChat = `INSERT INTO order_chat (invoice_no, created_by, updated_by, status) VALUES (?, ?, ?, ?)`

func (cr *chatRepo) InsertOrderChat(ctx context.Context, invoiceNo string, userID int64) (int64, error) {
	result, err := cr.db.GetMaster().ExecContext(ctx, insertOrderChat, invoiceNo, userID, userID, entity.ChatStatusOpen)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

const insertOrderChatHistory = `INSERT INTO order_chat_history (order_chat_id, content, origin, created_by, updated_by) VALUES (?, ?, ?, ?, ?)`

func (cr *chatRepo) InsertOrderChatHistory(ctx context.Context, data entity.OrderChatHistory) error {
	_, err := cr.db.GetMaster().ExecContext(ctx, insertOrderChatHistory, data.OrderChatID, data.Content, data.UserID, data.UserID, data.UserID)
	return err
}

const getOrderChatList = `SELECT content, origin, chat_type FROM order_chat_history WHERE order_chat_id = ? ORDER BY id ASC`

func (cr *chatRepo) GetOrderChatList(ctx context.Context, orderChatID int64) ([]entity.OrderChatHistoryList, error) {
	resp := []entity.OrderChatHistoryList{}
	err := cr.db.GetSlave().SelectContext(ctx, &resp, getOrderChatList, orderChatID)
	return resp, err
}
