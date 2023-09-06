package db

import (
	"context"

	"minitok/cmd/relation/global"
	"minitok/model"
)

func QueryLatestMessage(ctx context.Context, fromUid, toUid int64, query string) (*model.Message, error) {
	var message model.Message
	err := global.GormDB.WithContext(ctx).Select(query).Order("created_at desc").Where("from_user_id = ? AND to_user_id = ?", fromUid, toUid).First(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}
