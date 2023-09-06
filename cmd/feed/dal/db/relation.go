package db

import (
	"context"

	"minitok/cmd/feed/global"
	"minitok/model"
)

func QueryFollowInfo(ctx context.Context, userId, followUserId int64, query string) error {
	return global.GormDB.WithContext(ctx).Select(query).Where("user_id = ? AND follow_user_id = ?", userId, followUserId).First(&model.Relation{}).Error
}
