package entity

type UserRole int

const (
	UserSeller UserRole = iota + 1
	UserBuyer
)

type VerifiedStatus int

const (
	UserNotVerified VerifiedStatus = iota
	UserVerified
)

type UserInfo struct {
	ID          int64          `db:"id"`
	Email       string         `db:"email" json:"email"`
	Password    string         `db:"password" json:"password"`
	Name        string         `db:"name" json:"name"`
	PhoneNumber string         `db:"phone_number"`
	Avatar      string         `db:"avatar"`
	IsVerified  VerifiedStatus `db:"is_verified"`
}

type ResetPassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserToken struct {
	UserID      int64
	AccessToken string
	ExpiredAt   int64
}
