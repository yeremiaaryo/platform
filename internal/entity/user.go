package entity

type UserRole int

const (
	UserSeller UserRole = iota + 1
	UserBuyer
)

type UserInfo struct {
	ID   int64    `db:"id"`
	Name string   `db:"name"`
	Role UserRole `db:"role"`
}
