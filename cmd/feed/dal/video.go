package dal

import (
	"context"
	"minitok/cmd/feed/constant"
	"minitok/cmd/feed/dal/cache"
	"minitok/cmd/feed/dal/db"
	"minitok/cmd/feed/global"
	"minitok/model"
)

func QueryVideoInfoById(ctx context.Context, vid int64) (*model.Video, error) {
	videoInfo, err := cache.GetVideoInfo(ctx, vid)
	if err != nil {
		// 缓存未命中
		videoInfo, err = db.QueryVideoInfoById(ctx, vid, constant.VideoBaseInfoQueryString)
		if err != nil {
			return nil, err
		}
	}
	_ = cache.NewVideoInfos(ctx, []*model.Video{videoInfo}, global.ExpireDurationVideoBaseInfo)

	videoCounter, err := cache.GetVideoCounter(ctx, vid)
	if err != nil {
		videoCounter, err = db.QueryVideoInfoById(ctx, vid, constant.VideoCounterInfoQueryString)
		if err != nil {
			return nil, err
		}
		_ = cache.NewVideoCounters(ctx, []*model.Video{videoCounter})
	}
	videoInfo.FavoriteCount = videoCounter.FavoriteCount
	videoInfo.CommentCount = videoCounter.CommentCount
	return videoInfo, nil
}
