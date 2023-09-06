package db

import (
	"context"

	"minitok/cmd/comment/global"
	"minitok/model"
)

func QueryFirstUserInfoByID(ctx context.Context, id int64, query string) (*model.User, error) {
	var user model.User
	err := global.GormDB.WithContext(ctx).Select(query).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
