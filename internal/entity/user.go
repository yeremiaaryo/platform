package entity

type UserRole int

const (
	UserSeller UserRole = iota + 1
	UserBuyer
)

type UserInfo struct {
	ID          int64  `db:"id"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	Name        string `db:"name" json:"name"`
	PhoneNumber string `db:"phone_number"`
	Avatar      string `db:"avatar"`
}

type ResetPassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
