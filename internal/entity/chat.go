package entity

type ChatStatus int

const (
	ChatStatusOpen ChatStatus = iota + 1
	ChatStatusClosed
)

type OrderChat struct {
	ID        int64      `db:"id"`
	InvoiceNo string     `db:"invoice_no"`
	Status    ChatStatus `db:"status"`
}

type OrderChatHistory struct {
	OrderChatID int64
	Content     string
	UserID      int64
}

type OrderChatRequest struct {
	InvoiceNo string `json:"invoice_no"`
	Content   string `json:"content"`
}

type ChatType int

const (
	ChatTypeNormal ChatType = iota
	ChatTypeImage
	ChatTypeCatalog
)

type OrderChatHistoryList struct {
	ContentDB []byte      `db:"content" json:"-"`
	Content   interface{} `json:"content"`
	User      int64       `db:"origin" json:"user_id_origin"`
	ChatType  ChatType    `db:"chat_type" json:"chat_type"`
}
