package repository

import (
	"context"
	"database/sql"

	"github.com/yeremiaaryo/platform/internal/entity"
)

const fetchUser = `SELECT id, password, name, email FROM user WHERE email = ?`

func (ur *userRepo) FetchUserDataByEmail(ctx context.Context, email string) (*entity.UserInfo, error) {
	resp := new(entity.UserInfo)
	err := ur.db.GetSlave().GetContext(ctx, resp, fetchUser, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return resp, err
}

const registerUser = `INSERT INTO user (email, password, name) VALUES (?, ?, ?)`

func (ur *userRepo) RegisterUser(ctx context.Context, user entity.UserInfo) error {
	_, err := ur.db.GetMaster().ExecContext(ctx, registerUser, user.Email, user.Password, user.Name)
	return err
}
