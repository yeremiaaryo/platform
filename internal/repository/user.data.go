package repository

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

const fetchUser = `SELECT id, name, role FROM user WHERE id = ?`

func (ur *userRepo) FetchUserDataByUserID(ctx context.Context, userID int64) (*entity.UserInfo, error) {
	resp := new(entity.UserInfo)
	err := ur.db.GetSlave().GetContext(ctx, resp, fetchUser, userID)
	return resp, err
}
