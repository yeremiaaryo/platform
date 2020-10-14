package repository

import (
	"context"
	"database/sql"

	"github.com/yeremiaaryo/platform/internal/entity"
)

const fetchWaitinglist = `SELECT COUNT(1) FROM waitinglist WHERE email = ?`

func (ur *userRepo) FetchWaitinglist(ctx context.Context, email string) (int64, error) {
	resp := int64(0)
	err := ur.db.GetSlave().GetContext(ctx, resp, fetchWaitinglist, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return resp, err
}

const registerWaitinglist = `INSERT INTO waitinglist (email) VALUES (?)`

func (ur *userRepo) RegisterWaitinglist(ctx context.Context, email string) error {
	_, err := ur.db.GetMaster().ExecContext(ctx, registerWaitinglist, email)
	return err
}

const fetchUser = `SELECT id, password, name, email, is_verified FROM user WHERE email = ?`

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

func (ur *userRepo) RegisterUser(ctx context.Context, user entity.UserInfo) (int64, error) {
	result, err := ur.db.GetMaster().ExecContext(ctx, registerUser, user.Email, user.Password, user.Name)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

const updatePassword = `UPDATE user SET password = ? WHERE email = ?`

func (ur *userRepo) ResetPassword(ctx context.Context, data entity.ResetPassword, email string) error {
	_, err := ur.db.GetMaster().ExecContext(ctx, updatePassword, data.NewPassword, email)
	return err
}

const updateVerified = `UPDATE user SET is_verified = ? WHERE id = ?`

func (ur *userRepo) UpdateVerifiedUser(ctx context.Context, userID int64) error {
	_, err := ur.db.GetMaster().ExecContext(ctx, updateVerified, entity.UserVerified, userID)
	return err
}
