package usecase

import "context"

func (uu *userUC) GetUserName(ctx context.Context, userID int64) (string, error) {
	return uu.userSvc.GetUserName(ctx, userID)
}
