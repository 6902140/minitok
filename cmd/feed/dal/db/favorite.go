package db

import (
	"context"

	"minitok/cmd/feed/global"
	"minitok/model"
)

func QueryFavoriteInfo(ctx context.Context, uid, vid int64) error {
	return global.GormDB.WithContext(ctx).Select("id").Where("user_id = ? AND video_id = ?", uid, vid).First(&model.Favorite{}).Error
}
