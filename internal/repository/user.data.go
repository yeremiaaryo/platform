package repository

import (
	"context"

	"github.com/yeremiaaryo/platform/internal/entity"
)

const fetchUser = `SELECT id, name FROM user WHERE id = ?`

func (ur *userRepo) FetchUserDataByUserID(ctx context.Context, userID int64) (*entity.UserInfo, error) {
	resp := new(entity.UserInfo)
	err := ur.db.GetSlave().GetContext(ctx, resp, fetchUser, userID)
	return resp, err
}

const registerUser = `INSERT INTO user (email, password, name) VALUES (?, ?, ?)`

func (ur *userRepo) RegisterUser(ctx context.Context, user entity.UserInfo) error {
	_, err := ur.db.GetMaster().ExecContext(ctx, registerUser, user.Email, user.Password, user.Name)
	return err
}
