package service

import (
	"context"
	"fmt"
	"minitok/cmd/publish/dal"
	"minitok/cmd/publish/dal/cache"
	"minitok/cmd/publish/dal/mongodb"
	"minitok/cmd/publish/global"
	"minitok/kitex_gen/douyin/feed"
	"minitok/kitex_gen/douyin/publish"
	"minitok/kitex_gen/douyin/user"
	"minitok/model"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// PublishList 只需要点赞状态、点赞数以及封面截图
func (s *PublishListService) PublishList(req *publish.ListRequest) ([]*feed.Video, error) {
	var err error
	// 1. 空值缓存判定
	if err = cache.GetUserInfoNullKey(s.ctx, req.UserId); err == nil {
		return nil, errno.UserNotRegisterErr
	}
	if err = cache.GetPublishInfoNullKey(s.ctx, req.UserId); err == nil {
		return nil, nil
	}
	// 2. 查找发布的视频vid
	vidList, err := mongodb.GetPublishInfo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(vidList) == 0 {
		// 没有发布的视频, 添加空值缓存
		_ = cache.AddPublishInfoNullKey(s.ctx, req.UserId, global.ExpireDurationNullKey)
		return nil, nil
	}

	//3. 根据vid查找视频信息 缓存->db
	videoInfos := make([]*model.Video, len(vidList))
	for i, vid := range vidList {
		if videoInfos[i], err = dal.QueryVideoInfoById(s.ctx, vid); err != nil {
			return nil, err
		}
	}
	// 4. 解析Token
	var userId int64
	if len(req.Token) != 0 {
		claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
		if err == nil && claims.Id != 0 && claims.Issuer == global.Configs.JWT.Issuer && claims.Subject == global.Configs.JWT.Subject {
			userId = claims.Id
		}
	}
	// 5. model.Video -> feed.Video
	res := make([]*feed.Video, len(videoInfos))
	for i, videoInfo := range videoInfos {
		res[i] = new(feed.Video)
		res[i].Id = videoInfo.Id
		res[i].Title = videoInfo.Title
		res[i].Author = &user.User{Id: req.UserId}
		res[i].PlayUrl = fmt.Sprintf("%s/%s", global.Configs.FileAccess.NginxUrl, videoInfo.VideoPath)
		res[i].CoverUrl = fmt.Sprintf("%s/%s", global.Configs.FileAccess.NginxUrl, videoInfo.CoverPath)
		res[i].FavoriteCount = videoInfo.FavoriteCount
		res[i].CommentCount = videoInfo.CommentCount
		res[i].IsFavorite = false
		if userId != 0 {
			//if res[i].IsFavorite, err = cache.GetFavoriteStatus(s.ctx, userId, videoInfo.Id); err != nil {
			//	return nil, err
			//}
			if res[i].IsFavorite, err = mongodb.GetFavoriteInfo(s.ctx, userId, videoInfo.Id); err != nil {
				return nil, err
			}
		}
	}
	return res, nil
}
