package db

import (
	"context"

	"gorm.io/gorm"
	"minitok/cmd/publish/global"
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

func IncrPublishCountByUserId(ctx context.Context, id int64) error {
	return global.GormDB.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("work_count", gorm.Expr("price + 1")).Error
}
