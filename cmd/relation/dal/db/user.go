package db

import (
	"context"

	"minitok/cmd/relation/global"
	"minitok/model"
)

func QueryFirstUserInfoByUsername(ctx context.Context, username, query string) (*model.User, error) {
	var user model.User
	err := global.GormDB.WithContext(ctx).Select(query).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUserInfos(ctx context.Context, userInfos []*model.User) error {
	return global.GormDB.WithContext(ctx).Create(&userInfos).Error
}

func QueryUserInfoByID(ctx context.Context, id int64, query string) (*model.User, error) {
	var userInfo *model.User
	err := global.GormDB.WithContext(ctx).Select(query).Where("id = ?", id).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
